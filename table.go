package dynamock

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"foxygo.at/s/errs"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var ErrDuplicate = errors.New("duplicate")

type Table struct {
	Name     string    `json:"name"`
	Schema   Schema    `json:"schema"`
	RawItems []RawItem `json:"items"`

	items     []Item
	byPrimary map[string]map[string]Item   // lookup of Primary key by partition and sort key - unique result required.
	byGSI     map[string]map[string][]Item // lookup of globalSecondaryIndex by index name and partition key, sorted by sortKey
}

type RawItem = map[string]interface{} // as read from file
type Item = map[string]*dynamodb.AttributeValue

type Schema struct {
	PrimaryKey KeyDef   `json:"primaryKey"`
	GSIs       []KeyDef `json:"globalSecondaryIndex,omitempty"`
}

type KeyDef struct {
	Name         string      `json:"name,omitempty"`
	PartitionKey KeyPartDef  `json:"partitionKey"`
	SortKey      *KeyPartDef `json:"sortKey,omitempty"`
}

type KeyPartDef struct {
	Name string `json:"name"`
	Type string `json:"type"` // string, number;  binary  not implemented
}

func (t *Table) covertRawItems() error {
	t.items = make([]map[string]*dynamodb.AttributeValue, len(t.RawItems))
	for i, rawItem := range t.RawItems {
		item, err := dynamodbattribute.MarshalMap(rawItem)
		if err != nil {
			return errs.Errorf("covertRawItems: table %s marshalMap: %v", t.Name, err)
		}
		t.items[i] = item
	}
	return nil
}

func (t *Table) WriteSnap(w io.Writer, cols []string) error {
	if err := dynamodbattribute.UnmarshalListOfMaps(t.items, &t.RawItems); err != nil {
		return err
	}
	format := t.rowFormat(cols)
	row := make([]interface{}, len(cols))
	untypedCols := make([]interface{}, len(cols))
	for i, c := range cols {
		untypedCols[i] = c
	}
	fmt.Fprintf(w, format, untypedCols...)
	for _, rawItem := range t.RawItems {
		for i, col := range cols {
			row[i] = rawItem[col]
		}
		fmt.Fprintf(w, format, row...)
	}
	return nil
}

// Derived from rawItems, so rawItems must be set first
func (t *Table) rowFormat(cols []string) string {
	pads := make([]int, len(cols))
	for i, c := range cols {
		pads[i] = len(c)
	}
	for _, item := range t.RawItems {
		for i, c := range cols {
			attr := item[c]
			l := len(fmt.Sprint(attr))
			if l > pads[i] {
				pads[i] = l
			}
		}
	}
	formats := make([]string, len(cols))
	for i, pad := range pads {
		formats[i] = `%` + strconv.Itoa(pad) + `v`
	}
	return strings.Join(formats, ", ") + "\n"
}

// derived from items, must be set
func (t *Table) index() error {
	t.initByMaps()
	for _, item := range t.items {
		if err := t.indexItem(item); err != nil {
			return err
		}
	}
	return nil
}

func (t *Table) initByMaps() {
	t.byPrimary = map[string]map[string]Item{}
	if len(t.Schema.GSIs) == 0 {
		return
	}
	t.byGSI = map[string]map[string][]Item{}
	for _, gsi := range t.Schema.GSIs {
		t.byGSI[gsi.Name] = map[string][]Item{}
	}
}

func (t *Table) indexItem(item Item) error {
	if err := t.indexItemByPrimaryKey(item); err != nil {
		return err
	}
	for _, gsi := range t.Schema.GSIs {
		t.indexItemByGSI(item, gsi)
	}
	return nil
}

func (t *Table) indexItemByPrimaryKey(item Item) error {
	pk := t.Schema.PrimaryKey
	partKey := keyString(item, &pk.PartitionKey)
	sortKey := keyString(item, pk.SortKey)

	if t.byPrimary[partKey] == nil {
		t.byPrimary[partKey] = map[string]Item{}
	} else if _, ok := t.byPrimary[partKey][sortKey]; ok {
		return errs.Errorf("%v: %v: partionKey '%s', sortKey: '%s'", ErrDuplicate, ErrPrimaryKeyVal, partKey, sortKey)
	}
	t.byPrimary[partKey][sortKey] = item
	return nil
}

func (t *Table) indexItemByGSI(item Item, gsi KeyDef) {
	if !hasKey(item, gsi) {
		return
	}
	partKey := keyString(item, &gsi.PartitionKey)
	items := t.byGSI[gsi.Name][partKey]
	t.byGSI[gsi.Name][partKey] = insertItem(items, item, gsi.SortKey)
}

// append if no sortKey provided
// otherwise insert into sorted list of items at right position.
// sort order determined by sortKey
func insertItem(items []Item, item Item, sortKey *KeyPartDef) []Item {
	if sortKey == nil {
		return append(items, item)
	}
	less := lessFunc(items, item, sortKey)
	i := sort.Search(len(items), less)
	// insert at index i
	return append(items[:i], append([]Item{item}, items[i:]...)...)
}

// create comparison function for insertion in sorted slice of items ans
// use with sort.Search, based on key name and type.
func lessFunc(items []Item, item Item, key *KeyPartDef) func(int) bool {
	if key.Type == "string" {
		return func(i int) bool {
			return *items[i][key.Name].S >= *item[key.Name].S
		}
	}
	// key.Type: "number"
	f, _ := strconv.ParseFloat(*item[key.Name].N, 64)
	return func(i int) bool {
		fi, _ := strconv.ParseFloat(*items[i][key.Name].N, 64)
		return fi >= f
	}
}

func keyString(item Item, key *KeyPartDef) string {
	if key == nil || item[key.Name] == nil {
		return ""
	}
	if key.Type == "string" {
		return *item[key.Name].S
	}
	if key.Type == "number" {
		return *item[key.Name].N
	}
	return ""
}
