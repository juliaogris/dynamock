package dynamock

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"sync"

	"foxygo.at/s/errs"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var ErrDuplicate = errors.New("duplicate")

const primaryName = "/" // no name clashes possible as "/" is illegal for index names in dynamodb

type Table struct {
	m        sync.RWMutex
	Name     string    `json:"name"`
	Schema   Schema    `json:"schema"`
	RawItems []RawItem `json:"items"`

	items []Item
	// byPrimary is a lookup of Primary key by partition and sort key - unique result required.
	byPrimary map[string]map[string]Item
	// byIndex is a lookup by indexName and partition key.
	// indexName is either Global Secondary Index name or "/" for primaryKey.
	// the resulting slice of Items is sorted by sortKey
	byIndex map[string]map[string][]Item
}

type RawItem = map[string]interface{} // as read from file
type Item = map[string]*dynamodb.AttributeValue

func ItemToJSON(i Item) string {
	var out interface{}
	_ = dynamodbattribute.UnmarshalMap(i, &out)
	b, _ := json.Marshal(out)
	return string(b)
}

type Schema struct {
	PrimaryKey KeyDef   `json:"primaryKey"`
	GSIs       []KeyDef `json:"globalSecondaryIndex,omitempty"`
	gsis       map[string]KeyDef
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

func (t *Table) convertRawItems() error {
	t.m.Lock()
	defer t.m.Unlock()
	t.items = make([]map[string]*dynamodb.AttributeValue, len(t.RawItems))
	for i, rawItem := range t.RawItems {
		item, err := dynamodbattribute.MarshalMap(rawItem)
		if err != nil {
			return errs.Errorf("convertRawItems: table %s marshalMap: %v", t.Name, err)
		}
		t.items[i] = item
	}
	t.Schema.gsis = map[string]KeyDef{}
	for _, gsi := range t.Schema.GSIs {
		t.Schema.gsis[gsi.Name] = gsi
	}
	pk := t.Schema.PrimaryKey
	pk.Name = primaryName
	t.Schema.gsis[pk.Name] = pk
	return nil
}

func (t *Table) WriteSnap(w io.Writer, cols []string) error {
	t.m.RLock()
	defer t.m.RUnlock()
	return WriteSnap(w, t.items, cols)
}

func SnapString(items []Item, cols []string) string {
	sb := &bytes.Buffer{}
	_ = WriteSnap(sb, items, cols)
	return sb.String()
}

func WriteSnap(w io.Writer, items []Item, cols []string) error {
	iitems := []map[string]interface{}{}
	_ = dynamodbattribute.UnmarshalListOfMaps(items, &iitems)
	format := rowFormat(cols, iitems)
	row := make([]interface{}, len(cols))
	untypedCols := make([]interface{}, len(cols))
	for i, c := range cols {
		untypedCols[i] = c
	}
	fmt.Fprintf(w, format, untypedCols...)
	for _, rawItem := range iitems {
		for i, col := range cols {
			row[i] = rawItem[col]
		}
		fmt.Fprintf(w, format, row...)
	}
	return nil
}

// Derived from rawItems, so rawItems must be set first
func rowFormat(cols []string, iitems []map[string]interface{}) string {
	pads := make([]int, len(cols))
	for i, c := range cols {
		pads[i] = len(c)
	}
	for _, item := range iitems {
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
	t.byPrimary = map[string]map[string]Item{}
	t.byIndex = map[string]map[string][]Item{}
	for name := range t.Schema.gsis {
		t.byIndex[name] = map[string][]Item{}
	}
	for _, item := range t.items {
		if err := t.indexItem(item); err != nil {
			return err
		}
	}
	return nil
}

func (t *Table) indexItem(item Item) error {
	if err := t.indexItemByPrimaryKey(item); err != nil {
		return err
	}
	for _, gsi := range t.Schema.gsis {
		t.indexItemByKey(item, gsi)
	}
	return nil
}

func (t *Table) indexItemByPrimaryKey(item Item) error {
	pk := t.Schema.PrimaryKey
	k, _ := getKeyStrings(item, pk)
	if t.byPrimary[k.PartitionKey] == nil {
		t.byPrimary[k.PartitionKey] = map[string]Item{}
	} else if _, ok := t.byPrimary[k.PartitionKey][k.SortKey]; ok {
		return errs.Errorf("%v: %v: partionKey '%s', sortKey: '%s'", ErrDuplicate, ErrPrimaryKeyVal, k.PartitionKey, k.SortKey)
	}
	t.byPrimary[k.PartitionKey][k.SortKey] = item
	return nil
}

func (t *Table) indexItemByKey(item Item, gsi KeyDef) {
	if !hasKey(item, gsi) {
		return
	}
	k, _ := getKeyStrings(item, gsi)
	items := t.byIndex[gsi.Name][k.PartitionKey]
	t.byIndex[gsi.Name][k.PartitionKey] = insertItem(items, item, gsi.SortKey)
}

func (t *Table) Delete(key Item) (Item, error) {
	t.m.Lock()
	defer t.m.Unlock()
	if err := validateKeyItem(key, t.Schema); err != nil {
		return nil, err
	}
	return t.pop(key), nil
}

func (t *Table) Get(key Item) (Item, error) {
	t.m.RLock()
	defer t.m.RUnlock()
	if err := validateKeyItem(key, t.Schema); err != nil {
		return nil, err
	}
	k, _ := getKeyStrings(key, t.Schema.PrimaryKey)
	return t.get(k), nil
}

func (t *Table) Put(item Item) (Item, error) {
	t.m.Lock()
	defer t.m.Unlock()
	if err := validateItem(item, t.Schema); err != nil {
		return nil, err
	}
	old := t.pop(item)
	t.items = append(t.items, item)
	_ = t.indexItem(item)
	return old, nil
}

func (t *Table) Query(k *keyCondExpr, gsi *string, forward bool, exclusiveStartKey Item) ([]Item, error) {
	t.m.RLock()
	defer t.m.RUnlock()
	var items []Item
	index := primaryName
	key := t.Schema.PrimaryKey
	if gsi != nil {
		index = *gsi
		key = t.Schema.gsis[index]
	}
	if key.PartitionKey.Name != k.partitionCond.keyName {
		return nil, errs.Errorf("Query: KeyCondidtion: %v: want: %s got %s", ErrInvalidKey, t.Schema.PrimaryKey.PartitionKey.Name, k.partitionCond.keyName)
	}
	s, err := getKeyString(k.partitionCond.av, key.PartitionKey.Type)
	if err != nil {
		return nil, err
	}
	items = t.byIndex[index][s]
	if !forward {
		items = reverse(items)
	}
	items, err = sliceAfterStartKey(items, exclusiveStartKey, t.Schema.PrimaryKey)
	if err != nil {
		return nil, err
	}
	return applySortKeyCond(items, k.sortCond), nil
}

func applySortKeyCond(items []Item, k *keyCond) []Item {
	if items == nil || k == nil {
		return items
	}
	result := []Item{}
	for _, item := range items {
		if k.Check(item) {
			result = append(result, item)
		}
	}
	return result
}

func sliceAfterStartKey(items []Item, exclusiveStartKey Item, key KeyDef) ([]Item, error) {
	if exclusiveStartKey == nil {
		return items, nil
	}
	startKeys, err := getKeyStrings(exclusiveStartKey, key)
	if err != nil {
		return nil, err
	}
	for i, item := range items {
		keys, err := getKeyStrings(item, key)
		if err != nil {
			return nil, err
		}
		if *startKeys == *keys {
			return items[i+1:], nil
		}
	}
	return nil, nil
}

func reverse(items []Item) []Item {
	if len(items) < 2 {
		return items
	}
	items2 := make([]Item, len(items))
	l := len(items) - 1
	for i, item := range items {
		items2[l-i] = item
	}
	return items2
}

func (t *Table) getLastEvaluatedKey(items, pagedItems []Item) Item {
	if len(items) == len(pagedItems) {
		return nil
	}
	lastItem := pagedItems[len(pagedItems)-1]
	key := t.Schema.PrimaryKey
	result := Item{
		key.PartitionKey.Name: lastItem[key.PartitionKey.Name],
	}
	if key.SortKey != nil {
		result[key.SortKey.Name] = lastItem[key.SortKey.Name]
	}
	return result
}

func (t *Table) get(k *keyStrings) Item {
	if t.byPrimary[k.PartitionKey] == nil {
		return nil
	}
	return t.byPrimary[k.PartitionKey][k.SortKey]
}

// item needs to be validated with ValidateItem
func (t *Table) pop(key Item) Item {
	pk := t.Schema.PrimaryKey
	k, _ := getKeyStrings(key, pk)
	old := t.get(k)
	if old == nil {
		return nil
	}
	for _, gsi := range t.Schema.GSIs {
		if hasKey(old, gsi) {
			gsiKey, _ := getKeyStrings(old, gsi)
			items := t.byIndex[gsi.Name][gsiKey.PartitionKey]
			t.byIndex[gsi.Name][gsiKey.PartitionKey] = t.deleteItemInSlice(items, k)
		}
	}
	t.items = t.deleteItemInSlice(t.items, k)
	delete(t.byPrimary[k.PartitionKey], k.SortKey)
	return old
}

func (t *Table) deleteItemInSlice(items []Item, delKeys *keyStrings) []Item {
	pk := t.Schema.PrimaryKey
	for i, item := range items {
		k, _ := getKeyStrings(item, pk)
		if k.PartitionKey == delKeys.PartitionKey && k.SortKey == delKeys.SortKey {
			return append(items[:i], items[i+1:]...)
		}
	}
	return items
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

func getKeyString(attr *dynamodb.AttributeValue, attrType string) (string, error) {
	if attr == nil {
		return "", errs.New(ErrInvalidKey, ErrNil)
	}
	switch attrType {
	case "string":
		if attr.S == nil {
			return "", errs.Errorf("%v: %v: %v, expected string", ErrInvalidKey, ErrInvalidType, attr)
		}
		return *attr.S, nil
	case "number":
		if attr.N == nil {
			return "", errs.Errorf("%v: %v: %v, expected number", ErrInvalidKey, ErrInvalidType, attr)
		}
		return *attr.N, nil
	}
	return "", errs.Errorf("%v: %v: %s expected 'string' or 'number'", ErrInvalidKey, ErrInvalidType, attr)
}

type keyStrings struct {
	PartitionKey string
	SortKey      string
}

func getKeyStrings(item Item, keyDef KeyDef) (*keyStrings, error) {
	partKey := keyDef.PartitionKey
	partKeyVal, err := getKeyString(item[partKey.Name], partKey.Type)
	if err != nil {
		return nil, err
	}
	sortKey := keyDef.SortKey
	sortKeyVal := ""
	if keyDef.SortKey != nil {
		sortKeyVal, err = getKeyString(item[sortKey.Name], sortKey.Type)
		if err != nil {
			return nil, err
		}
	}
	return &keyStrings{PartitionKey: partKeyVal, SortKey: sortKeyVal}, nil
}

func pageItems(items []Item, limit *int64, pageSize int) []Item {
	if limit == nil && pageSize == 0 {
		return items
	}
	l := len(items)
	if pageSize != 0 && pageSize < l {
		l = pageSize
	}
	if limit != nil && int(*limit) < l {
		l = int(*limit)
	}
	return items[:l]
}
