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

	tableNames []string
	tables     map[string]*Table
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
		if err := t.convertRawItems(); err != nil {
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

func (db *DB) GetItem(in *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	if in == nil {
		return nil, errs.Errorf("GetItem: %v: GetItemInput", ErrNil)
	}
	if in.AttributesToGet != nil || in.ExpressionAttributeNames != nil || in.ProjectionExpression != nil {
		msg := "GetItemInput fields: AttributesToGet, ExpressionAttributeNames, ProjectionExpression, ReturnConsumedCapacity"
		return nil, errs.Errorf("GetItem: %v: %s", ErrUnimpl, msg)
	}
	if err := validateTableName(db, in.TableName); err != nil {
		return nil, err
	}
	table := db.tables[*in.TableName]
	item, err := table.Get(in.Key)
	if err != nil {
		return nil, err
	}
	return &dynamodb.GetItemOutput{Item: item}, nil
}

func (db *DB) GetItemWithContext(_ aws.Context, in *dynamodb.GetItemInput, _ ...request.Option) (*dynamodb.GetItemOutput, error) {
	return db.GetItem(in)
}

func (db *DB) PutItem(in *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	if in == nil {
		return nil, errs.Errorf("%v: PutItemInput", ErrNil)
	}
	if in.ConditionExpression != nil || in.ConditionalOperator != nil ||
		in.Expected != nil || in.ExpressionAttributeNames != nil ||
		in.ExpressionAttributeValues != nil {
		msg := "ConditionExpression, ConditionalOperator, Expected, ExpressionAttributeNames, ExpressionAttributeValues, ReturnConsumedCapacity, ReturnItemCollectionMetrics, ReturnValues"
		return nil, errs.Errorf("PutItem: %v: %s", ErrUnimpl, msg)
	}
	if err := validateTableName(db, in.TableName); err != nil {
		return nil, err
	}
	table := db.tables[*in.TableName]
	old, err := table.Put(in.Item)
	if err != nil {
		return nil, err
	}
	if in.ReturnValues != nil && *in.ReturnValues == "ALL_OLD" {
		return &dynamodb.PutItemOutput{Attributes: old}, nil
	}
	return &dynamodb.PutItemOutput{Attributes: old}, nil
}

func (db *DB) PutItemWithContext(_ aws.Context, in *dynamodb.PutItemInput, _ ...request.Option) (*dynamodb.PutItemOutput, error) {
	return db.PutItem(in)
}

func (db *DB) DeleteItem(in *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	if in == nil {
		return nil, errs.Errorf("%v: DeleteItemInput", ErrNil)
	}
	if in.ConditionExpression != nil || in.ConditionalOperator != nil ||
		in.Expected != nil || in.ExpressionAttributeNames != nil ||
		in.ExpressionAttributeValues != nil {
		msg := "ConditionExpression, ConditionalOperator, Expected, ExpressionAttributeNames, ExpressionAttributeValues, ReturnConsumedCapacity, ReturnItemCollectionMetrics, ReturnValues"
		return nil, errs.Errorf("DeleteItem: %v: %s", ErrUnimpl, msg)
	}
	if err := validateTableName(db, in.TableName); err != nil {
		return nil, err
	}
	table := db.tables[*in.TableName]
	old, err := table.Delete(in.Key)
	if err != nil {
		return nil, err
	}
	if in.ReturnValues != nil && *in.ReturnValues == "ALL_OLD" {
		return &dynamodb.DeleteItemOutput{Attributes: old}, nil
	}
	return &dynamodb.DeleteItemOutput{Attributes: old}, nil
}

func (db *DB) DeleteItemWithContext(_ aws.Context, in *dynamodb.DeleteItemInput, _ ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	return db.DeleteItem(in)
}
