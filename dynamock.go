// Package dynamock provides a file-based dynamoDB fake called DB.
//
// DB implements the dynamodbiface.DynamoDBAPI interface.
// Use it for testing, local development and execution.
// Dynamodb docs:
//   https://pkg.go.dev/github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface
//   https://pkg.go.dev/github.com/aws/aws-sdk-go/service/dynamodb
package dynamock

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"foxygo.at/s/errs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Table struct {
	Name     string    `json:"name"`
	Schema   Schema    `json:"schema"`
	RawItems []RawItem `json:"items"`

	items []Item
	// byPrimaryKey      map[string]map[string]Item   // lookup of Primary key by partition and sort key
	// byGSIKey          map[string]map[string][]Item // lookup of globalSecondaryIndex by index name and partition key, sorted by sortKey
	// byGSIPartitionKey map[string]map[string][]Item // lookup of globalSecondaryIndex by index name and partition key, sorted by sortKey
}

type RawItem = map[string]interface{} // *dynamodb.AttributeValue
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
	Type string `json:"type"` // string, number, binary
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

// Derived from rawItems!!
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

// DB implements the dynamodbiface.DynamoDBAPI interface
type DB struct {
	UnimplementedDB

	RawTables []*Table `json:"tables"`

	tables map[string]*Table
}

func NewDB() *DB {
	return &DB{}
}

func NewDBFromReader(r io.Reader) (*DB, error) {
	db := &DB{tables: map[string]*Table{}}
	if err := json.NewDecoder(r).Decode(db); err != nil {
		return nil, errs.Errorf("NewDBFromReader: %v", err)
	}
	for _, t := range db.RawTables {
		if err := t.covertRawItems(); err != nil {
			return nil, err
		}
	}
	if err := validateDB(db); err != nil {
		return nil, err
	}
	for _, t := range db.RawTables {
		db.tables[t.Name] = t
	}
	db.index()
	return db, nil
}

func (db *DB) index() {
	//for _, table := range db.Tables {
	//fmt.Println(item)
	//}
}

func (db *DB) WriteSnap(w io.Writer) error {
	for _, t := range db.tables {
		if err := dynamodbattribute.UnmarshalListOfMaps(t.items, &t.RawItems); err != nil {
			return err
		}
	}
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	return e.Encode(db)
}

func (db *DB) GetItem(_ *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return nil, ErrUnimpl
}

func (db *DB) GetItemWithContext(_ aws.Context, _ *dynamodb.GetItemInput, _ ...request.Option) (*dynamodb.GetItemOutput, error) {
	return nil, ErrUnimpl
}
