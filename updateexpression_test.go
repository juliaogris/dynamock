package dynamock

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitWords(t *testing.T) {
	words := []string{" SET ", " REMOVE "}
	got := splitWords("a SET b REMOVE c SET REMOVE", words)
	want := []string{"a", " SET b", " REMOVE c", " SET REMOVE"}
	require.Equal(t, want, got)

	got = splitWords("a SET b REMOVE c SET REMOVE ", words)
	want = []string{"a", " SET b", " REMOVE c", " SET", " REMOVE "}
	require.Equal(t, want, got)

	got = splitWords("abc", words)
	want = []string{"abc"}
	require.Equal(t, want, got)

	got = splitWords("", words)
	want = []string{""}
	require.Equal(t, want, got)

	got = splitWords("SET", words)
	want = []string{"SET"}
	require.Equal(t, want, got)

	got = splitWords(" SET ", words)
	want = []string{" SET "}
	require.Equal(t, want, got)

	got = splitWords(" SET x", words)
	want = []string{" SET x"}
	require.Equal(t, want, got)

	got = splitWords("x SET ", words)
	want = []string{"x", " SET "}
	require.Equal(t, want, got)
}

func TestIndexWords(t *testing.T) {
	words := []string{" SET ", " REMOVE "}
	require.Equal(t, 3, indexWords("abc SET 123 REMOVE", words, -1))
	require.Equal(t, 11, indexWords("abc sET 123 REMOVE ", words, 0))
	require.Equal(t, -1, indexWords("abc sET 123 REMOVE", words, 0))
	require.Equal(t, 0, indexWords(" SET 123 REMOVE", words, -1))
	require.Equal(t, -1, indexWords(" SET 123 REMOVE", words, 0))
}

func TestTrimSpaceWords(t *testing.T) {
	words := []string{" SET ", " REMOVE", "abc ", "xyz"}
	want := []string{"SET", "REMOVE", "abc", "xyz"}
	require.Equal(t, want, trimSpaceWords(words))
}

func TestParseUpdateExprErr(t *testing.T) {
	_, err := parseUpdateExpr(nil, nil, nil)
	requireErrIs(t, err, ErrNil)

	_, err = parseUpdateExpr(strPtr(" BAD_UPDATE_EXPR "), nil, nil)
	requireErrIs(t, err, ErrInvalidUpdateExpression)

	_, err = parseUpdateExpr(strPtr(" SET bad-set-expr "), nil, nil)
	requireErrIs(t, err, ErrInvalidUpdateExpression)

	_, err = parseUpdateExpr(strPtr(" REMOVE //bad-set-expr "), nil, nil)
	requireErrIs(t, err, ErrInvalidUpdateExpression)

	_, err = parseUpdateExpr(strPtr(" REMOVE #attr"), nil, nil)
	requireErrIs(t, err, ErrSubstitution)

	_, err = parseUpdateExpr(strPtr("SET //bad_attr=:name"), nil, nil)
	requireErrIs(t, err, ErrInvalidUpdateExpression)

	_, err = parseUpdateExpr(strPtr("SET name=missing_colon"), nil, nil)
	requireErrIs(t, err, ErrInvalidUpdateExpression)

	_, err = parseUpdateExpr(strPtr("SET #name=:name"), nil, nil)
	requireErrIs(t, err, ErrSubstitution)

	_, err = parseUpdateExpr(strPtr("SET name=:name"), nil, nil)
	requireErrIs(t, err, ErrSubstitution)
}
