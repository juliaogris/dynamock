package dynamock

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/require"
)

func TestParseKeyCondExpr(t *testing.T) {
	s := strPtr("id = :id")
	valueSub := Item{":id": {S: strPtr("123")}}
	got, err := parseKeyCondExpr(s, valueSub, nil)
	require.NoError(t, err)
	want := &keyCondExpr{
		partitionCond: keyCond{
			keyName: "id",
			op:      eq,
			val:     ":id",
			av:      valueSub[":id"],
		},
	}
	require.Equal(t, want, got)

	s = strPtr("id=:id")
	got, err = parseKeyCondExpr(s, valueSub, nil)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestParseKeyCondExprWithSort(t *testing.T) {
	s := strPtr("#name = :name AND age > :age")
	valueSub := Item{
		":name": {S: strPtr("Joe")},
		":age":  {N: strPtr("18")},
	}
	nameSub := map[string]*string{"#name": strPtr("name")}
	got, err := parseKeyCondExpr(s, valueSub, nameSub)
	require.NoError(t, err)
	want := &keyCondExpr{
		partitionCond: keyCond{keyName: "name", op: eq, val: ":name", av: valueSub[":name"]},
		sortCond:      &keyCond{keyName: "age", op: greater, val: ":age", av: valueSub[":age"]},
	}
	require.Equal(t, want, got)

	s = strPtr("name=:name AND age>:age")
	got, err = parseKeyCondExpr(s, valueSub, nil)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestParseKeyCondExprPartitionErr(t *testing.T) {
	s := strPtr("id<:id")
	valueSub := Item{":id": {S: strPtr("123")}}
	_, err := parseKeyCondExpr(s, valueSub, nil)
	requireErrIs(t, err, ErrInvalidKeyCondition)
	s = strPtr("id")
	_, err = parseKeyCondExpr(s, nil, nil)
	requireErrIs(t, err, ErrInvalidKeyCondition)
}

func TestOpString(t *testing.T) {
	require.Equal(t, "=", eq.String())
	require.Equal(t, ">", greater.String())
	require.Equal(t, "<", less.String())
	require.Equal(t, ">=", greaterEq.String())
	require.Equal(t, "<=", lessEq.String())
	require.Equal(t, "BETWEEN", between.String())
	require.Equal(t, "begins_with", beginsWith.String())
	require.Equal(t, "UNKNOWN", op(255).String())
}

func TestKeyCondCheck(t *testing.T) {
	av := &dynamodb.AttributeValue{S: strPtr("1")}
	k := keyCond{op: eq, keyName: "id", av: av}
	require.True(t, k.Check(Item{"id": av}))
	require.False(t, k.Check(Item{"id2": av}))
	require.False(t, k.Check(Item{"id": &dynamodb.AttributeValue{SS: []*string{}}}))

	k = keyCond{op: greaterEq, keyName: "id", av: av}
	require.True(t, k.Check(Item{"id": av}))

	k = keyCond{op: lessEq, keyName: "id", av: av}
	require.True(t, k.Check(Item{"id": av}))

	k = keyCond{op: greater, keyName: "id", av: av}
	require.False(t, k.Check(Item{"id": av}))

	k = keyCond{op: less, keyName: "id", av: av}
	require.False(t, k.Check(Item{"id": av}))

	k = keyCond{op: between, keyName: "id", av: av}
	require.False(t, k.Check(Item{"id": av}))

	av2 := &dynamodb.AttributeValue{S: strPtr("2")}
	k = keyCond{op: between, keyName: "id", av: av, av2: av2}
	require.True(t, k.Check(Item{"id": av}))

	avN := &dynamodb.AttributeValue{N: strPtr("1")}
	k = keyCond{op: lessEq, keyName: "id", av: avN}
	require.True(t, k.Check(Item{"id": avN}))

	k = keyCond{op: less, keyName: "id", av: avN}
	require.False(t, k.Check(Item{"id": avN}))

	k = keyCond{op: between, keyName: "id", av: avN}
	require.False(t, k.Check(Item{"id": avN}))

	avN2 := &dynamodb.AttributeValue{N: strPtr("2")}
	k = keyCond{op: between, keyName: "id", av: avN, av2: avN2}
	require.True(t, k.Check(Item{"id": avN}))

	k = keyCond{op: eq, keyName: "id", av: avN}
	require.False(t, k.Check(Item{"id": av}))

	k = keyCond{op: eq, keyName: "id", av: av}
	require.False(t, k.Check(Item{"id": avN}))

	k = keyCond{op: op(255), keyName: "id", av: av}
	require.False(t, k.Check(Item{"id": av}))

	k = keyCond{op: op(255), keyName: "id", av: avN}
	require.False(t, k.Check(Item{"id": avN}))
}

func TestParseKeyCondExprStrErr(t *testing.T) {
	_, err := parseKeyCondExprStr(nil)
	requireErrIs(t, err, ErrInvalidKeyCondition)

	e := "x = :x AND y BAD_OP :y"
	_, err = parseKeyCondExprStr(&e)
	requireErrIs(t, err, ErrInvalidKeyCondition)

	e = "x = :x AND begins_with(y, :y"
	_, err = parseKeyCondExprStr(&e)
	requireErrIs(t, err, ErrInvalidKeyCondition)
}

func TestParseKeyCondExprErr(t *testing.T) {
	e := "x = :x"
	_, err := parseKeyCondExpr(&e, nil, nil)
	requireErrIs(t, err, ErrInvalidKeyCondition)

	e = "x = :x AND y <= :y"
	_, err = parseKeyCondExpr(&e, Item{":x": &dynamodb.AttributeValue{S: strPtr("123")}}, nil)
	requireErrIs(t, err, ErrInvalidKeyCondition)

	e = "x = :x AND y BETWEEN :y1 AND :y2"
	subs := Item{
		":x":  &dynamodb.AttributeValue{S: strPtr("123")},
		":y1": &dynamodb.AttributeValue{S: strPtr("123")},
	}
	_, err = parseKeyCondExpr(&e, subs, nil)
	requireErrIs(t, err, ErrInvalidKeyCondition)

	e = "#x = :x"
	subs = Item{":x": &dynamodb.AttributeValue{S: strPtr("123")}}
	_, err = parseKeyCondExpr(&e, subs, nil)
	requireErrIs(t, err, ErrSubstitution)

	nameSubs := map[string]*string{"#BAD_X": strPtr("x")}
	_, err = parseKeyCondExpr(&e, subs, nameSubs)
	requireErrIs(t, err, ErrSubstitution)
}

func TestNewKeyCond(t *testing.T) {
	_, err := newKeyCond(nil, eq)
	requireErrIs(t, err, ErrInvalidKeyCondition)

	s := []string{"///INVALID_CHARS", ":y"}
	_, err = newKeyCond(s, eq)
	requireErrIs(t, err, ErrInvalidKeyCondition)

	s = []string{"x", "///INVALID_CHARS"}
	_, err = newKeyCond(s, eq)
	requireErrIs(t, err, ErrInvalidKeyCondition)

	s = []string{"x", ":y1", "///INVALID_CHARS"}
	_, err = newKeyCond(s, between)
	requireErrIs(t, err, ErrInvalidKeyCondition)
}
