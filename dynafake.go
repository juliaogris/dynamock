// Package dynafake provides a file-based dynamoDB fake called DB.
//
// DB implements the dynamodbiface.DynamoDBAPI interface.
// Use it for testing, local development and execution.
// Dynamodb docs:
//   https://pkg.go.dev/github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface
//   https://pkg.go.dev/github.com/aws/aws-sdk-go/service/dynamodb
package dynafake

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

// DB implements the dynamodbiface.DynamoDBAPI interface
type DB struct {
	dynamodbiface.DynamoDBAPI
}

func NewDB() *DB {
	return &DB{}
}
