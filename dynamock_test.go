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
	require.Truef(t, errors.Is(err, target), "expected target: '%v' got: '%v'", target, err)
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

func strPtr(s string) *string {
	return &s
}

func CloseIgnoreErr(c io.Closer) {
	_ = c.Close()
}

func TestDynamoDBAPI(t *testing.T) {
	iface := (*dynamodbiface.DynamoDBAPI)(nil)
	require.Implements(t, iface, NewDB())
	require.Implements(t, iface, &UnimplementedDB{})
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

func TestWriteDBSnap(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")

	sb := &bytes.Buffer{}
	err := db.WriteSnap(sb)
	require.NoError(t, err)

	want := string(ReadTestdataBytes(t, "db.json"))
	require.JSONEq(t, want, sb.String())
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
				TableName: strPtr("person"),
				Key:       Item{"id": {N: strPtr("0")}},
			},
			want: `{ "id": 0, "name": "Jon", "phone": "000", "age": 0 }`,
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

func TestGetItemUnimplErr(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	in := &dynamodb.GetItemInput{
		TableName:       strPtr("product"),
		AttributesToGet: []*string{strPtr("?")},
	}
	_, err := db.GetItem(in)
	requireErrIs(t, err, ErrUnimpl)
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

func TestGetKeyString(t *testing.T) {
	attr := &dynamodb.AttributeValue{N: strPtr("12")}
	got, err := getKeyString(attr, "number")
	require.NoError(t, err)
	require.Equal(t, "12", got)
}

func TestGetKeyStringErr(t *testing.T) {
	attr := &dynamodb.AttributeValue{N: strPtr("12")}
	_, err := getKeyString(attr, "string")
	require.Error(t, err)
	requireErrIs(t, err, ErrInvalidType)

	attr = &dynamodb.AttributeValue{S: strPtr("abc")}
	_, err = getKeyString(attr, "number")
	require.Error(t, err)
	requireErrIs(t, err, ErrInvalidType)

	_, err = getKeyString(attr, "badType")
	require.Error(t, err)
	requireErrIs(t, err, ErrInvalidType)
}

func TestGetKeyStringsErr(t *testing.T) {
	key := Item{}
	keyDef := KeyDef{PartitionKey: KeyPartDef{Name: "id", Type: "string"}}
	_, err := getKeyStrings(key, keyDef)
	requireErrIs(t, err, ErrInvalidKey)

	key = Item{
		"id0": {S: strPtr("0")},
		"id1": {S: strPtr("1")},
		"id2": {S: strPtr("2")},
		"id3": {S: strPtr("3")},
	}
	_, err = getKeyStrings(key, keyDef)
	requireErrIs(t, err, ErrInvalidKey)
}

func TestPutItem(t *testing.T) {
	testCases := map[string]struct {
		inTableName    string
		inItem         Item
		inKey          Item
		wantLenDelta   int
		wantAttributes Item
	}{
		"new_product": {
			inTableName:    "product",
			inKey:          Item{"id": {S: strPtr("100")}},
			inItem:         Item{"id": {S: strPtr("100")}, "name": {S: strPtr("sticky notes")}, "price": {N: strPtr("1")}},
			wantLenDelta:   1,
			wantAttributes: nil,
		},
		"replace_product": {
			inTableName:    "product",
			inKey:          Item{"id": {S: strPtr("1")}},
			inItem:         Item{"id": {S: strPtr("1")}, "name": {S: strPtr("sticky notes")}, "price": {N: strPtr("1")}},
			wantLenDelta:   0,
			wantAttributes: Item{"id": {S: strPtr("1")}, "name": {S: strPtr("red pen")}, "price": {N: strPtr("11")}},
		},
		"new_path": {
			inTableName:    "path",
			inKey:          Item{"folder": {S: strPtr("/Users/dev/")}, "file": {S: strPtr("README.md")}},
			inItem:         Item{"folder": {S: strPtr("/Users/dev/")}, "file": {S: strPtr("README.md")}, "perms": {S: strPtr("-rw-r--r--")}},
			wantLenDelta:   1,
			wantAttributes: nil,
		},
		"replace_path": {
			inTableName:    "path",
			inKey:          Item{"folder": {S: strPtr("/Users/dev/")}, "file": {S: strPtr("todo.txt")}},
			inItem:         Item{"folder": {S: strPtr("/Users/dev/")}, "file": {S: strPtr("todo.txt")}, "perms": {S: strPtr("-r--r--r--")}},
			wantLenDelta:   0,
			wantAttributes: Item{"folder": {S: strPtr("/Users/dev/")}, "file": {S: strPtr("todo.txt")}, "perms": {S: strPtr("-rw-r--r--")}},
		},
	}
	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			db := ReadTestdataDB(t, "db.json")
			length := len(db.tables[tc.inTableName].items)
			in := &dynamodb.PutItemInput{
				TableName:    &tc.inTableName,
				Item:         tc.inItem,
				ReturnValues: strPtr("ALL_OLD"),
			}
			out, err := db.PutItem(in)
			require.NoError(t, err)
			require.Equal(t, tc.wantAttributes, out.Attributes)
			wantLength := length + tc.wantLenDelta
			require.Equal(t, wantLength, len(db.tables[tc.inTableName].items))

			in2 := &dynamodb.GetItemInput{
				TableName: &tc.inTableName,
				Key:       tc.inKey,
			}
			out2, err := db.GetItem(in2)
			require.NoError(t, err)
			require.Equal(t, tc.inItem, out2.Item)
		})
	}
}

func lenGSI(gsi map[string][]Item) int {
	n := 0
	for _, v := range gsi {
		n = n + len(v)
	}
	return n
}

func lenPrimary(byPrimary map[string]map[string]Item) int {
	n := 0
	for _, v := range byPrimary {
		n = n + len(v)
	}
	return n
}

func TestPutNewItemIndexUpdate(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	in := &dynamodb.PutItemInput{
		TableName: strPtr("person"),
		Item: Item{
			"id":   {N: strPtr("100")},
			"name": {S: strPtr("Hector")},
			"age":  {N: strPtr("100")},
		},
	}
	length := len(db.tables["person"].items)
	lengthPrimary := lenPrimary(db.tables["person"].byPrimary)
	lengthPhoneGSI := lenGSI(db.tables["person"].byGSI["phoneGSI"])
	lengthNameGSI := lenGSI(db.tables["person"].byGSI["nameGSI"])
	out, err := db.PutItem(in)
	require.NoError(t, err)
	require.Nil(t, out.Attributes)
	require.Equal(t, length+1, len(db.tables["person"].items))
	require.Equal(t, lengthPrimary+1, lenPrimary(db.tables["person"].byPrimary))
	require.Equal(t, lengthPhoneGSI, lenGSI(db.tables["person"].byGSI["phoneGSI"]))
	require.Equal(t, lengthNameGSI+1, lenGSI(db.tables["person"].byGSI["nameGSI"]))
}

func TestPutReplaceItemIndexUpdate(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	in := &dynamodb.PutItemInput{
		TableName: strPtr("person"),
		Item: Item{
			"id":    {N: strPtr("0")},
			"name":  {S: strPtr("Hector")},
			"phone": {S: strPtr("1001000")},
		},
		ReturnValues: strPtr("ALL_OLD"),
	}
	length := len(db.tables["person"].items)
	lengthPrimary := lenPrimary(db.tables["person"].byPrimary)
	lengthPhoneGSI := lenGSI(db.tables["person"].byGSI["phoneGSI"])
	lengthNameGSI := lenGSI(db.tables["person"].byGSI["nameGSI"])
	out, err := db.PutItem(in)
	require.NoError(t, err)
	want := Item{"id": {N: strPtr("0")}, "name": {S: strPtr("Jon")}, "phone": {S: strPtr("000")}, "age": {N: strPtr("0")}}
	require.Equal(t, want, out.Attributes)
	require.Equal(t, length, len(db.tables["person"].items))
	require.Equal(t, lengthPrimary, lenPrimary(db.tables["person"].byPrimary))
	require.Equal(t, lengthPhoneGSI, lenGSI(db.tables["person"].byGSI["phoneGSI"]))
	require.Equal(t, lengthNameGSI-1, lenGSI(db.tables["person"].byGSI["nameGSI"]))
}

func TestPutErr(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	_, err := db.PutItem(nil)
	requireErrIs(t, err, ErrNil)

	in := &dynamodb.PutItemInput{
		ConditionExpression: strPtr("??"),
	}
	_, err = db.PutItem(in)
	requireErrIs(t, err, ErrUnimpl)

	in = &dynamodb.PutItemInput{
		TableName: strPtr("bad_table_name"),
	}
	_, err = db.PutItem(in)
	requireErrIs(t, err, ErrUnknownTable)

	_, err = db.PutItemWithContext(context.Background(), in)
	requireErrIs(t, err, ErrUnknownTable)
}

func TestPutInvalidItemErr(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	in := &dynamodb.PutItemInput{
		TableName: strPtr("person"),
		Item:      nil,
	}
	_, err := db.PutItem(in)
	requireErrIs(t, err, ErrMissingAttribute)
}

func TestDeleteItem(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	in := &dynamodb.DeleteItemInput{
		TableName: strPtr("person"),
		Key: Item{
			"id": {N: strPtr("1")},
		},
		ReturnValues: strPtr("ALL_OLD"),
	}
	length := len(db.tables["person"].items)
	lengthPrimary := lenPrimary(db.tables["person"].byPrimary)
	lengthPhoneGSI := lenGSI(db.tables["person"].byGSI["phoneGSI"])
	lengthNameGSI := lenGSI(db.tables["person"].byGSI["nameGSI"])
	out, err := db.DeleteItem(in)
	require.NoError(t, err)
	want := `{ "id": 1, "name": "Jon", "phone": "111", "age": 11 }`
	got := ItemToJSON(out.Attributes)
	require.JSONEq(t, want, got)
	require.Equal(t, length-1, len(db.tables["person"].items))
	require.Equal(t, lengthPrimary-1, lenPrimary(db.tables["person"].byPrimary))
	require.Equal(t, lengthPhoneGSI-1, lenGSI(db.tables["person"].byGSI["phoneGSI"]))
	require.Equal(t, lengthNameGSI-1, lenGSI(db.tables["person"].byGSI["nameGSI"]))

	// delete again
	out, err = db.DeleteItem(in)
	require.NoError(t, err)
	require.Nil(t, out.Attributes)
	require.Equal(t, length-1, len(db.tables["person"].items))
	require.Equal(t, lengthPrimary-1, lenPrimary(db.tables["person"].byPrimary))
	require.Equal(t, lengthPhoneGSI-1, lenGSI(db.tables["person"].byGSI["phoneGSI"]))
	require.Equal(t, lengthNameGSI-1, lenGSI(db.tables["person"].byGSI["nameGSI"]))

	// and again
	in.ReturnValues = nil
	out, err = db.DeleteItem(in)
	require.NoError(t, err)
	require.Nil(t, out.Attributes)
	require.Equal(t, length-1, len(db.tables["person"].items))
	require.Equal(t, lengthPrimary-1, lenPrimary(db.tables["person"].byPrimary))
	require.Equal(t, lengthPhoneGSI-1, lenGSI(db.tables["person"].byGSI["phoneGSI"]))
	require.Equal(t, lengthNameGSI-1, lenGSI(db.tables["person"].byGSI["nameGSI"]))
}

func TestDeleteItemErr(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	_, err := db.DeleteItem(nil)
	requireErrIs(t, err, ErrNil)

	in := &dynamodb.DeleteItemInput{
		ConditionExpression: strPtr("??"),
	}
	_, err = db.DeleteItem(in)
	requireErrIs(t, err, ErrUnimpl)

	in = &dynamodb.DeleteItemInput{
		TableName: strPtr("bad_table_name"),
	}
	_, err = db.DeleteItem(in)
	requireErrIs(t, err, ErrUnknownTable)

	_, err = db.DeleteItemWithContext(context.Background(), in)
	requireErrIs(t, err, ErrUnknownTable)
}

func TestDeleteInvalidItemErr(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	in := &dynamodb.DeleteItemInput{
		TableName: strPtr("person"),
		Key:       nil,
	}
	_, err := db.DeleteItem(in)
	requireErrIs(t, err, ErrInvalidKey)

	_, err = db.DeleteItemWithContext(context.Background(), in)
	requireErrIs(t, err, ErrInvalidKey)
}

func TestDeleteItemInSlice(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	out := db.tables["product"].deleteItemInSlice(nil, nil, nil)
	require.Nil(t, out)
}
