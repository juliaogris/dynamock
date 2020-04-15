// Package dynafake contains a file-based dynamoDB fake and utilities.
// DB implements the dynamodbiface.DynamoDBAPI interface
// https://pkg.go.dev/github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface
// It is useful for testing, local development and execution.
package dynafake

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

// DB implements the dynamodbiface.DynamoDBAPI interface
type DB struct {
	dynamodbiface.DynamoDBAPI
}

func NewDB() *DB {
	return &DB{}
}
