package dynafake

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
)

func TestDynamoDBAPI(t *testing.T) {
	iface := (*dynamodbiface.DynamoDBAPI)(nil)
	assert.Implements(t, iface, NewDB())
}
