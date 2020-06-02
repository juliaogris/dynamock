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

type JSONTable struct {
	Name   string                   `json:"name"`
	Schema Schema                   `json:"schema"`
	Items  []map[string]interface{} `json:"items"`
}

type JSONDB struct {
	Tables []*JSONTable `json:"tables"`
}

// DB implements the dynamodbiface.DynamoDBAPI interface
type DB struct {
	UnimplementedDB

	tableNames []string
	tables     map[string]*Table
	pageSize   int
}

func NewDB() *DB {
	return &DB{}
}

func NewDBFromReader(r io.Reader) (*DB, error) {
	jdb := JSONDB{}
	if err := json.NewDecoder(r).Decode(&jdb); err != nil {
		return nil, errs.Errorf("NewDBFromReader: %v", err)
	}
	db := &DB{tables: map[string]*Table{}}
	for _, t := range jdb.Tables {
		items := make([]Item, len(t.Items))
		for i, item := range t.Items {
			av, _ := dynamodbattribute.MarshalMap(item)
			items[i] = av
		}
		table := &Table{name: t.Name, schema: t.Schema, items: items}
		if err := validateTable(table); err != nil {
			return nil, err
		}
		if err := table.index(); err != nil {
			return nil, err
		}
		db.tableNames = append(db.tableNames, table.name)
		db.tables[table.name] = table
	}
	return db, nil
}

func (db *DB) WriteSnap(w io.Writer) error {
	jdb := JSONDB{
		Tables: make([]*JSONTable, len(db.tables)),
	}
	for i, name := range db.tableNames {
		t := db.tables[name]
		items := []map[string]interface{}{}
		_ = dynamodbattribute.UnmarshalListOfMaps(t.items, &items)
		jdb.Tables[i] = &JSONTable{Schema: t.schema, Name: t.name, Items: items}
	}
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	return e.Encode(jdb)
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

func (db *DB) Query(in *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	if err := validateQueryIntput(in); err != nil {
		return nil, err
	}
	if err := validateTableName(db, in.TableName); err != nil {
		return nil, err
	}
	table := db.tables[*in.TableName]
	if err := validateIndexName(table, in.IndexName); err != nil {
		return nil, err
	}
	keyCond, err := parseKeyCondExpr(in.KeyConditionExpression, in.ExpressionAttributeValues, in.ExpressionAttributeNames)
	if err != nil {
		return nil, err
	}
	forward := true
	if in.ScanIndexForward != nil && !*in.ScanIndexForward {
		forward = false
	}
	items, err := table.Query(keyCond, in.IndexName, forward, in.ExclusiveStartKey)
	if err != nil {
		return nil, err
	}
	pagedItems := pageItems(items, in.Limit, db.pageSize)
	if in.Select != nil && *in.Select == "COUNT" {
		count := int64(len(pagedItems))
		return &dynamodb.QueryOutput{Count: &count, ScannedCount: &count}, nil
	}
	out := &dynamodb.QueryOutput{
		Items:            pagedItems,
		LastEvaluatedKey: table.getLastEvaluatedKey(items, pagedItems),
	}
	return out, nil
}

func validateQueryIntput(in *dynamodb.QueryInput) error {
	if in == nil {
		return errs.Errorf("%v: QueryInput", ErrNil)
	}
	if in.AttributesToGet != nil || in.ConditionalOperator != nil ||
		in.FilterExpression != nil || in.KeyConditions != nil ||
		in.ProjectionExpression != nil || in.QueryFilter != nil {
		msg := "AttributesToGet, ConditionalOperator, FilterExpression, KeyConditions, ProjectionExpression, QueryFilter"
		return errs.Errorf("QueryItem: %v: %s", ErrUnimpl, msg)
	}
	if in.Select != nil && (*in.Select == "SPECIFIC_ATTRIBUTES" || *in.Select == "ALL_PROJECTED_ATTRIBUTES") {
		return errs.Errorf("QueryItem: %v: %s", ErrUnimpl, *in.Select)
	}
	if in.KeyConditionExpression == nil {
		return errs.Errorf("%v: KeyConditionExpression", ErrNil)
	}
	if in.ExpressionAttributeValues == nil {
		return errs.Errorf("%v: missing ExpressionAttributeValues", ErrSubstitution)
	}
	return nil
}

func (db *DB) QueryWithContext(_ aws.Context, in *dynamodb.QueryInput, _ ...request.Option) (*dynamodb.QueryOutput, error) {
	return db.Query(in)
}

func (db *DB) UpdateItem(in *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	if in == nil || in.UpdateExpression == nil || in.Key == nil || in.ExpressionAttributeValues == nil {
		return nil, errs.Errorf("%v: UpdateItemInput [UpdateExpression | Key | ExpressionAttributeValues]", ErrNil)
	}
	if in.AttributeUpdates != nil || in.ConditionalOperator != nil || in.Expected != nil {
		msg := "AttributeUpdates, ConditionalOperator, Expected"
		return nil, errs.Errorf("UpdateItemInput: %v: %s", ErrUnimpl, msg)
	}
	if in.ReturnValues != nil && *in.ReturnValues != "NONE" && *in.ReturnValues != "ALL_OLD" && *in.ReturnValues != "ALL_NEW" {
		return nil, errs.Errorf("UpdateItemInput.ReturnValues: %v: expected NONE, ALL_OLD or ALL_NEW", ErrUnimpl)
	}
	if err := validateTableName(db, in.TableName); err != nil {
		return nil, err
	}
	table := db.tables[*in.TableName]
	updateExpr, err := parseUpdateExpr(in.UpdateExpression, in.ExpressionAttributeValues, in.ExpressionAttributeNames)
	if err != nil {
		return nil, err
	}
	item, err := table.Update(in.Key, updateExpr, in.ReturnValues)
	if err != nil {
		return nil, err
	}
	return &dynamodb.UpdateItemOutput{Attributes: item}, nil
}

func (db *DB) UpdateItemWithContext(_ aws.Context, in *dynamodb.UpdateItemInput, _ ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	return db.UpdateItem(in)
}
