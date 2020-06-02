package dynamock

import (
	"strings"

	"foxygo.at/s/errs"
)

var ErrInvalidUpdateExpression = errs.Errorf("invalid update expression")

type updateExpr struct {
	setExpr          Item
	removeAttributes []string
}

// Limited to SET a=:value1, b=:value2 .... REMOVE c, d, e // at this stage
func parseUpdateExpr(str *string, valueSub Item, nameSub map[string]*string) (*updateExpr, error) {
	if str == nil {
		return nil, errs.Errorf("%v: %v", ErrNil, ErrInvalidUpdateExpression)
	}
	s := strings.TrimSpace(*str)
	if !hasPrefixWords(s, []string{"SET ", "REMOVE "}) {
		return nil, errs.Errorf("%v: bad prefix '%s'", ErrInvalidUpdateExpression, s)
	}
	expr := trimSpaceWords(splitWords(s, []string{" SET ", " REMOVE "}))
	setExpr := Item{}
	var removeAttr []string
	for _, e := range expr {
		if strings.HasPrefix(e, "SET ") {
			e = strings.TrimSpace(strings.TrimPrefix(e, "SET "))
			if err := addSetExpr(e, setExpr, valueSub, nameSub); err != nil {
				return nil, err
			}
		} else if strings.HasPrefix(e, "REMOVE ") {
			var err error
			e = strings.TrimSpace(strings.TrimPrefix(e, "REMOVE "))
			if removeAttr, err = addRemoveAttr(e, removeAttr, nameSub); err != nil {
				return nil, err
			}
		}
	}
	return &updateExpr{setExpr: setExpr, removeAttributes: removeAttr}, nil
}

func addSetExpr(es string, setExpr Item, valueSub Item, nameSub map[string]*string) error {
	pairs := strings.Split(es, ",")
	for _, pair := range pairs {
		p := strings.Split(pair, "=")
		if len(p) != 2 {
			return errs.Errorf("%v: expected one '=' in '%s'", ErrInvalidUpdateExpression, pair)
		}
		name := strings.TrimSpace(p[0])
		val := strings.TrimSpace(p[1])
		if !reExprName.MatchString(name) {
			return errs.Errorf("%v: invalid expression attribute name '%s' ", ErrInvalidUpdateExpression, name)
		}
		if !reExprVal.MatchString(val) {
			return errs.Errorf("%v: invalid expression attribute value '%s' ", ErrInvalidUpdateExpression, val)
		}
		name2, err := substituteName(name, nameSub)
		if err != nil {
			return err
		}
		av, ok := valueSub[val]
		if !ok {
			return errs.Errorf("%v: %s", ErrSubstitution, val)
		}
		setExpr[name2] = av
	}
	return nil
}

func addRemoveAttr(es string, removeAttr []string, nameSub map[string]*string) ([]string, error) {
	attrs := strings.Split(es, ",")
	for _, attr := range attrs {
		a := strings.TrimSpace(attr)
		if !reExprName.MatchString(a) {
			return nil, errs.Errorf("%v: invalid expression attribute name '%s' ", ErrInvalidUpdateExpression, a)
		}
		a2, err := substituteName(a, nameSub)
		if err != nil {
			return nil, err
		}
		if !containsStr(removeAttr, a2) {
			removeAttr = append(removeAttr, a2)
		}
	}
	return removeAttr, nil
}

func containsStr(ss []string, str string) bool {
	for _, s := range ss {
		if s == str {
			return true
		}
	}
	return false
}

func hasPrefixWords(s string, words []string) bool {
	for _, word := range words {
		if strings.HasPrefix(s, word) {
			return true
		}
	}
	return false
}

func splitWords(s string, words []string) []string {
	var result []string
	for i := indexWords(s, words, 0); i != -1; i = indexWords(s, words, 0) {
		result = append(result, s[:i])
		s = s[i:]
	}
	return append(result, s)
}

func indexWords(s string, words []string, afterIdx int) int {
	i := len(s)
	for _, word := range words {
		idx := strings.Index(s, word)
		if idx != -1 && idx < i && idx > afterIdx {
			i = idx
		}
	}
	if i == len(s) {
		return -1
	}
	return i
}

func trimSpaceWords(words []string) []string {
	s := make([]string, len(words))
	for i, word := range words {
		s[i] = strings.TrimSpace(word)
	}
	return s
}
