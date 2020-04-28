package dynamock

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/require"
)

func TestDynamoDBAPI(t *testing.T) {
	iface := (*dynamodbiface.DynamoDBAPI)(nil)
	require.Implements(t, iface, NewDB())
	require.Implements(t, iface, &UnimplementedDB{})
}

func strPtr(s string) *string {
	return &s
}

func CloseIgnoreErr(c io.Closer) {
	_ = c.Close()
}

func TestDBFromReader(t *testing.T) {
	fpath := filepath.Join("testdata", "db.json")
	r, err := os.Open(fpath)
	defer CloseIgnoreErr(r)
	require.NoError(t, err)

	db, err := NewDBFromReader(r)
	require.NoError(t, err)
	require.NotNil(t, db)

	// first table schema
	require.Equal(t, 2, len(db.RawTables))
	pt := db.RawTables[0]
	require.Equal(t, "product", pt.Name)
	pk := pt.Schema.PrimaryKey.PartitionKey
	require.Equal(t, "id", pk.Name)
	require.Equal(t, "string", pk.Type)
	sk := pt.Schema.PrimaryKey.SortKey
	require.Nil(t, sk)
	require.Equal(t, 0, len(pt.Schema.GSIs))
	require.Nil(t, pt.Schema.GSIs)

	// first table data
	require.Equal(t, 4, len(pt.items))
	got := pt.items[0]["id"]
	want := &dynamodb.AttributeValue{S: strPtr("1")}
	require.Equal(t, want, got)
	got = pt.items[2]["price"]
	want = &dynamodb.AttributeValue{N: strPtr("33")}
	require.Equal(t, want, got)

	// second table schema
	pt = db.RawTables[1]
	require.Equal(t, "person", pt.Name)
	pk = pt.Schema.PrimaryKey.PartitionKey
	require.Equal(t, "id", pk.Name)
	require.Equal(t, "number", pk.Type)
	sk = pt.Schema.PrimaryKey.SortKey
	require.Nil(t, sk)
	require.Equal(t, 2, len(pt.Schema.GSIs))
	wantS := KeyDef{
		Name:         "phoneGSI",
		PartitionKey: KeyPartDef{Name: "phone", Type: "string"},
		SortKey:      &KeyPartDef{Name: "name", Type: "string"},
	}
	gotS := pt.Schema.GSIs[1]
	require.Equal(t, wantS, gotS)

	// second table data
	require.Equal(t, 9, len(pt.items))
}

func TestDBFromReaderJSONErr(t *testing.T) {
	r := strings.NewReader(`{ "tables": ["truncated...`)
	_, err := NewDBFromReader(r)
	require.Error(t, err, err)
	require.True(t, errors.Is(err, io.ErrUnexpectedEOF))
}

func TestDBFromReaderValidateErr(t *testing.T) {
	r := strings.NewReader(`{"tables" : [ {} ] }`)
	_, err := NewDBFromReader(r)
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrMissingName))
}

func TestDBFromReaderDuplicateErr(t *testing.T) {
	r := strings.NewReader(`{"tables" : [
		{
			"name": "product",
			"schema": {
				"primaryKey": {
					"partitionKey": { "name": "id", "type": "string" }
				}
			},
			"items": [
				{ "id": "1", "name": "red pen", "price": 11 },
				{ "id": "1", "name": "blue pen", "price": 22 }
			]
		}
	 ] }`)
	_, err := NewDBFromReader(r)
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrDuplicate))
}

func TestGetItem(t *testing.T) { //nolint:funlen
	db := DB{}
	ctx := context.Background()

	_, err := db.GetItem(nil)
	requireErrUnimpl(t, err)

	_, err = db.GetItemWithContext(ctx, nil)
	requireErrUnimpl(t, err)
}

func TestWriteDBSnap(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")

	sb := &bytes.Buffer{}
	err := db.WriteSnap(sb)
	require.NoError(t, err)

	want := string(ReadTestdataBytes(t, "db.json"))
	require.JSONEq(t, want, sb.String())
}

func ReadTestdataBytes(t *testing.T, filename string) []byte {
	t.Helper()
	fpath := filepath.Join("testdata", filename)
	r, err := os.Open(fpath)
	require.NoError(t, err)
	defer r.Close()
	b, err := ioutil.ReadAll(r)
	require.NoError(t, err)
	return b
}

func ReadTestdataDB(t *testing.T, filename string) *DB {
	t.Helper()
	b := ReadTestdataBytes(t, filename)

	db, err := NewDBFromReader(bytes.NewReader(b))
	require.NoError(t, err)
	require.NotNil(t, db)
	return db
}
