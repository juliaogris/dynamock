package dynamock

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/require"
)

func idDef() KeyDef {
	return KeyDef{PartitionKey: KeyPartDef{Name: "id", Type: "string"}}
}

func TestValidate(t *testing.T) {
	r, err := os.Open(filepath.Join("testdata", "db.json"))
	require.NoError(t, err)

	db, err := NewDBFromReader(r)
	require.NoError(t, err)
	require.NoError(t, validateDB(db))
}

func TestValidateDBErr(t *testing.T) {
	db := &DB{
		RawTables: []*Table{{Name: ""}},
	}
	err := validateDB(db)
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrMissingName))
}

func TestValidateTableErr(t *testing.T) {
	tbl := &Table{Name: "table1"}
	err := validateTable(tbl)
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrMissingName))
}

func TestValidateTableGSIErr(t *testing.T) {
	tbl := &Table{
		Name: "table1",
		Schema: Schema{
			PrimaryKey: idDef(),
			GSIs: []KeyDef{
				{PartitionKey: KeyPartDef{}},
			},
		},
	}
	err := validateTable(tbl)
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrMissingName))
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
		Name: "table1",
		Schema: Schema{
			PrimaryKey: idDef(),
			GSIs:       []KeyDef{emailGSI},
		},
	}
	err := validateTable(tbl)
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrUnknownType), err)
}

func TestValidateTableItemErr(t *testing.T) {
	tbl := &Table{
		Name:   "table1",
		Schema: Schema{PrimaryKey: idDef()},
		items: []Item{
			{"name": &dynamodb.AttributeValue{S: strPtr("Joe")}},
		},
	}
	err := validateTable(tbl)
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrPrimaryKeyVal), err)
}

func TestValidateAttrKeyTypeErr(t *testing.T) {
	err := validateAttrKeyType(nil, "string")
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrMissingAttribute), err)

	err = validateAttrKeyType(&dynamodb.AttributeValue{}, "bad_type")
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrMissingType), err)
}

func TestValidateKey(t *testing.T) {
	key := idDef()
	item := Item{}
	err := validateKey(item, key)
	require.Error(t, err, err)
	require.True(t, errors.Is(err, ErrMissingAttribute))
}

func TestValidateItemGSIErr(t *testing.T) {
	schema := &Schema{
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
	require.True(t, errors.Is(err, ErrGSIVal))
}
