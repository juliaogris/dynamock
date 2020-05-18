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

func requireErrIs(t *testing.T, err, target error) {
	t.Helper()
	require.Error(t, err)
	require.Error(t, target)
	require.Truef(t, errors.Is(err, target), "expected err '%v' to be '%v'", err, target)
}

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
	require.Equal(t, 3, len(db.RawTables))
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

func TestGetItem(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")

	testCases := map[string]struct {
		in   *dynamodb.GetItemInput
		want string
	}{
		"string_key": {
			in: &dynamodb.GetItemInput{
				TableName: strPtr("product"),
				Key:       Item{"id": {S: strPtr("1")}},
			},
			want: `{ "id": "1", "name": "red pen", "price": 11 }`,
		},
		"numeric_key": {
			in: &dynamodb.GetItemInput{
				TableName: strPtr("product"),
				Key:       Item{"id": {S: strPtr("1")}},
			},
			want: `{ "id": "1", "name": "red pen", "price": 11 }`,
		},
		"composite_key": {
			in: &dynamodb.GetItemInput{
				TableName: strPtr("path"),
				Key: Item{
					"folder": {S: strPtr("/Users/dev/")},
					"file":   {S: strPtr("todo.txt")},
				},
			},
			want: `{ "folder": "/Users/dev/", "file": "todo.txt", "perms": "-rw-r--r--" }`,
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			out, err := db.GetItem(tc.in)
			require.NoError(t, err)
			require.NotNil(t, out)
			require.NotNil(t, out.Item)
			require.JSONEq(t, tc.want, ItemToJSON(out.Item))

			out, err = db.GetItemWithContext(context.Background(), tc.in)
			require.NoError(t, err)
			require.JSONEq(t, tc.want, ItemToJSON(out.Item))
		})
	}
}

func TestGetNoItem(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	in := &dynamodb.GetItemInput{
		TableName: strPtr("product"),
		Key:       Item{"id": {S: strPtr("-1")}},
	}
	out, err := db.GetItem(in)
	require.NoError(t, err)
	require.NotNil(t, out)
	require.Nil(t, out.Item)
}

func TestGetItemNilErr(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	_, err := db.GetItem(nil)
	requireErrIs(t, err, ErrNil)
}

func TestGetItemTableNameErr(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	in := &dynamodb.GetItemInput{
		TableName: nil,
	}
	_, err := db.GetItem(in)
	requireErrIs(t, err, ErrNil)
	in = &dynamodb.GetItemInput{
		TableName: strPtr("bad_table_name"),
	}
	_, err = db.GetItem(in)
	requireErrIs(t, err, ErrUnknownTable)
}

func TestGetItemKeyErr(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	in := &dynamodb.GetItemInput{
		TableName: strPtr("path"),
		Key: Item{
			"folderXXX": {S: strPtr("/Users/dev/")},
		},
	}
	_, err := db.GetItem(in)
	requireErrIs(t, err, ErrMissingAttribute)

	in = &dynamodb.GetItemInput{
		TableName: strPtr("path"),
		Key: Item{
			"folder":   {S: strPtr("/Users/dev/")},
			"fileXXXX": {S: strPtr("todo.txt")},
		},
	}
	_, err = db.GetItem(in)
	requireErrIs(t, err, ErrMissingAttribute)
}

func TestGetKeyVal(t *testing.T) {
	attr := &dynamodb.AttributeValue{N: strPtr("12")}
	got, err := getKeyVal(attr, "number")
	require.NoError(t, err)
	require.Equal(t, "12", got)
}

func TestGetKeyValErr(t *testing.T) {
	attr := &dynamodb.AttributeValue{N: strPtr("12")}
	_, err := getKeyVal(attr, "string")
	require.Error(t, err)
	requireErrIs(t, err, ErrInvalidType)

	attr = &dynamodb.AttributeValue{S: strPtr("abc")}
	_, err = getKeyVal(attr, "number")
	require.Error(t, err)
	requireErrIs(t, err, ErrInvalidType)

	_, err = getKeyVal(attr, "badType")
	require.Error(t, err)
	requireErrIs(t, err, ErrInvalidType)
}

func TestGetKeyValsErr(t *testing.T) {
	key := Item{}
	keyDef := KeyDef{PartitionKey: KeyPartDef{Name: "id", Type: "string"}}
	_, err := getKeyVals(key, keyDef)
	requireErrIs(t, err, ErrInvalidKey)

	key = Item{
		"id0": {S: strPtr("0")},
		"id1": {S: strPtr("1")},
		"id2": {S: strPtr("2")},
		"id3": {S: strPtr("3")},
	}
	_, err = getKeyVals(key, keyDef)
	requireErrIs(t, err, ErrInvalidKey)
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
