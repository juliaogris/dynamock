package dynamock

import (
	"bytes"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/stretchr/testify/require"
)

func itemFixture() Item {
	return Item{
		"id":   &dynamodb.AttributeValue{N: strPtr("1")},
		"name": &dynamodb.AttributeValue{S: strPtr("Grace")},
	}
}

func TestWriteTableSnap(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")

	productTable := db.tables["product"]
	require.NotNil(t, productTable)

	sb := &bytes.Buffer{}
	cols := []string{"id", "price"}
	err := productTable.WriteSnap(sb, cols)
	require.NoError(t, err)
	want := `
  id, price
   1,    11
   2,    22
   3,    33
1234,  1234
`[1:]
	require.Equal(t, want, sb.String())
}

func TestCovertRawItemsNoErr(t *testing.T) {
	tbl := Table{
		RawItems: []map[string]interface{}{
			{"id": make(chan int)},
		},
	}
	err := tbl.covertRawItems()
	// Should be an error imo, but dynamo silently ignores
	// types it cannot marshal. I couldn't work out how to make
	// it error.
	require.NoError(t, err)
}

func TestCovertRawItemsErr(t *testing.T) {
	names := struct {
		Names []*string `dynamodbav:",stringset"`
	}{
		Names: []*string{nil}, // nil value in stringset causes InvalidMarshalError
	}
	tbl := Table{
		RawItems: []map[string]interface{}{
			{"names": names},
		},
	}
	err := tbl.covertRawItems()
	require.Error(t, err)
	errInv := &dynamodbattribute.InvalidMarshalError{}
	require.True(t, errors.As(err, &errInv))
}

func TestIndexProductTable(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	err := db.index()
	require.NoError(t, err)
	pt := db.tables["product"]
	require.NotNil(t, pt)
	require.Equal(t, 0, len(pt.byGSI))
	require.Equal(t, 4, len(pt.byPrimary))
	ids := []string{"1", "2", "3", "1234"}
	for i, id := range ids {
		want := pt.items[i]
		got := pt.byPrimary[id][""]
		require.Equal(t, want, got)
	}
}

func TestIndexPersonTable(t *testing.T) {
	db := ReadTestdataDB(t, "db.json")
	pt := db.tables["person"]
	require.NotNil(t, pt)
	require.Equal(t, 2, len(pt.byGSI))
	require.Equal(t, 9, len(pt.byPrimary))
	// testdata/db.json - items
	// 	id: 0, name: "Jon",     phone: "000", age: 0
	// 	id: 1, name: "Jon",     phone: "111", age: 11
	// 	id: 2, name: "Tom",     phone: "222", age: 22
	// 	id: 3, name: "Bee",     phone: "333", age: 33
	// 	id: 4, name: "Jen",     phone: "444", age: 44
	// 	id: 5, name: "Jen",     phone: "555"
	// 	id: 6, name: "No-phone",              age: 1
	// 	id: 7, name: "No-age", phone: "777"
	// 	id: 8, name: "Jen",    phone: "222", age: 15
	ids := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8"}
	items := pt.items
	for i, id := range ids {
		want := items[i]
		got := pt.byPrimary[id][""]
		require.Equal(t, want, got)
	}

	nameGSI := pt.byGSI["nameGSI"]
	want := []Item{items[0], items[1]}
	got := nameGSI["Jon"]
	require.Equal(t, want, got)

	want = []Item{items[2]}
	got = nameGSI["Tom"]
	require.Equal(t, want, got)

	want = []Item{items[8], items[4]}
	got = nameGSI["Jen"]
	require.Equal(t, want, got)

	require.Nil(t, nameGSI["No-age"])

	phoneGSI := pt.byGSI["phoneGSI"]
	want = []Item{items[0]}
	got = phoneGSI["000"]
	require.Equal(t, want, got)

	want = []Item{items[8], items[2]}
	got = phoneGSI["222"]
	require.Equal(t, want, got)

	require.Nil(t, phoneGSI["No-phone"])
}

func TestKeyString(t *testing.T) {
	require.Equal(t, "", keyString(nil, nil))
	key := &KeyPartDef{
		Type: "bad type",
		Name: "name",
	}
	require.Equal(t, "", keyString(itemFixture(), key))
}

func TestInsertItem(t *testing.T) {
	item := itemFixture()
	items := insertItem(nil, item, nil)
	got := insertItem(items, item, nil)

	want := []Item{item, item}
	require.Equal(t, want, got)
}
