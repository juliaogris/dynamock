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
	"io"

	"foxygo.at/s/errs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

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
	if err := db.index(); err != nil {
		return nil, err
	}
	return db, nil
}

func (db *DB) index() error {
	for _, table := range db.tables {
		if err := table.index(); err != nil {
			return err
		}
	}
	return nil
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
