package dynamock

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/require"
)

func idDef() KeyDef {
	return KeyDef{PartitionKey: KeyPartDef{Name: "id", Type: "string"}}
}

func TestValidateTableErr(t *testing.T) {
	tbl := &Table{name: "table1"}
	err := validateTable(tbl)

	require.Error(t, err)
	require.True(t, errors.Is(err, ErrMissingName))
}

func TestValidateTableGSIErr(t *testing.T) {
	tbl := &Table{
		name: "table1",
		schema: Schema{
			PrimaryKey: idDef(),
			GSIs: []KeyDef{
				{PartitionKey: KeyPartDef{}},
			},
		},
	}
	err := validateTable(tbl)
	requireErrIs(t, err, ErrMissingName)
}

func TestValidateTableGSITypeErr(t *testing.T) {
	emailGSI := KeyDef{
		Name: "emailGSI",
		PartitionKey: KeyPartDef{
			Name: "email",
			Type: "typo",
		},
	}
	tbl := &Table{
		name: "table1",
		schema: Schema{
			PrimaryKey: idDef(),
			GSIs:       []KeyDef{emailGSI},
		},
	}
	err := validateTable(tbl)
	requireErrIs(t, err, ErrUnknownType)
}

func TestValidateTableItemErr(t *testing.T) {
	tbl := &Table{
		name:   "table1",
		schema: Schema{PrimaryKey: idDef()},
		items: []Item{
			{"name": &dynamodb.AttributeValue{S: strPtr("Joe")}},
		},
	}
	err := validateTable(tbl)
	requireErrIs(t, err, ErrPrimaryKeyVal)
}

func TestValidateAttrKeyTypeErr(t *testing.T) {
	err := validateAttrKeyType(nil, "string")
	requireErrIs(t, err, ErrMissingAttribute)

	err = validateAttrKeyType(&dynamodb.AttributeValue{}, "bad_type")
	requireErrIs(t, err, ErrMissingType)
}

func TestValidateKey(t *testing.T) {
	key := idDef()
	item := Item{}
	err := validateKey(item, key)
	requireErrIs(t, err, ErrMissingAttribute)
}

func TestValidateItemGSIErr(t *testing.T) {
	schema := Schema{
		PrimaryKey: idDef(),

		GSIs: []KeyDef{
			{
				PartitionKey: KeyPartDef{
					Name: "phone",
					Type: "number",
				},
			},
		},
	}
	item := Item{
		"id":    &dynamodb.AttributeValue{S: strPtr("1")},
		"phone": &dynamodb.AttributeValue{S: strPtr("0123-4566")}, // number expected
	}
	err := validateItem(item, schema)
	require.Error(t, err)
	requireErrIs(t, err, ErrGSIVal)
}

func TestValidateKeyItemErr(t *testing.T) {
	err := validateKeyItem(nil, Schema{})
	requireErrIs(t, err, ErrInvalidKey)

	item := Item{"id1": nil, "id2": nil, "id3": nil}
	err = validateKeyItem(item, Schema{})
	requireErrIs(t, err, ErrInvalidKey)
}

func TestValidateIndexName(t *testing.T) {
	tbl := &Table{byIndex: map[string]map[string][]Item{}}
	err := validateIndexName(tbl, strPtr("missing_index"))
	requireErrIs(t, err, ErrUnknownIndex)
}
