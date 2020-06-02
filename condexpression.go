package dynamock

import (
	"regexp"
	"strconv"
	"strings"

	"foxygo.at/s/errs"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	ErrInvalidKeyCondition = errs.Errorf("invalid key condition")
	ErrSubstitution        = errs.Errorf("missing substitution")

	reExprName = regexp.MustCompile(`^#?[0-9A-Za-z_-]+$`)
	reExprVal  = regexp.MustCompile(`^:[0-9A-Za-z_-]+$`)
)

type op int

const (
	eq op = iota
	greater
	less
	greaterEq
	lessEq
	between
	beginsWith
)

func (o op) String() string {
	switch o {
	case eq:
		return "="
	case greater:
		return ">"
	case less:
		return "<"
	case greaterEq:
		return ">="
	case lessEq:
		return "<="
	case between:
		return "BETWEEN"
	case beginsWith:
		return "begins_with"
	}
	return "UNKNOWN"
}

type keyCond struct {
	keyName string
	op      op
	val     string
	val2    string
	av      *dynamodb.AttributeValue
	av2     *dynamodb.AttributeValue
}

type keyCondExpr struct {
	partitionCond keyCond
	sortCond      *keyCond
}

func (k *keyCond) checkN(av *dynamodb.AttributeValue) bool {
	if av == nil || k.av == nil || av.N == nil || k.av.N == nil {
		return false
	}
	f1, _ := strconv.ParseFloat(*av.N, 64)
	f2, _ := strconv.ParseFloat(*k.av.N, 64)
	switch k.op {
	case eq:
		return f1 == f2
	case greater:
		return f1 > f2
	case less:
		return f1 < f2
	case greaterEq:
		return f1 >= f2
	case lessEq:
		return f1 <= f2
	case between:
		if k.av2 == nil || k.av2.N == nil {
			return false
		}
		f3, _ := strconv.ParseFloat(*k.av2.N, 64)
		return f2 <= f1 && f1 <= f3
	}
	return false
}

func (k *keyCond) checkS(av *dynamodb.AttributeValue) bool {
	if av == nil || k.av == nil || av.S == nil || k.av.S == nil {
		return false
	}
	s1 := *av.S
	s2 := *k.av.S
	switch k.op {
	case eq:
		return s1 == s2
	case greater:
		return s1 > s2
	case less:
		return s1 < s2
	case greaterEq:
		return s1 >= s2
	case lessEq:
		return s1 <= s2
	case between:
		if k.av2 == nil || k.av2.S == nil {
			return false
		}
		return s2 <= s1 && s1 <= *k.av2.S
	case beginsWith:
		return strings.HasPrefix(s1, s2)
	}
	return false
}

func (k *keyCond) Check(item Item) bool {
	av := item[k.keyName]
	if av == nil {
		return false
	}
	if av.S != nil {
		return k.checkS(av)
	}
	if av.N != nil {
		return k.checkN(av)
	}
	return false
}

// Valid KeyConditionsExpressions can take one of the following forms:
// partitionKeyName = :partitionkeyval
// partitionKeyName = :partitionkeyval AND sortKeyName <OP> :sortkeyval
//     <OP> → = > < >= <=
// partitionKeyName = :partitionkeyval AND sortKeyName BETWEEN :sortkeyval1 AND :sortkeyval2
// partitionKeyName = :partitionkeyval AND begins_with ( sortKeyName, :sortkeyval )
// valueSub: ExpressionAttributeValues
// nameSub: ExpressionAttributeNames
func parseKeyCondExpr(s *string, valueSub Item, nameSub map[string]*string) (*keyCondExpr, error) {
	kc, err := parseKeyCondExprStr(s)
	if err != nil {
		return nil, err
	}
	k, err := substituteKeyCondExpr(kc.partitionCond, valueSub, nameSub)
	if err != nil {
		return nil, err
	}
	kc.partitionCond = *k
	if kc.sortCond != nil {
		k, err := substituteKeyCondExpr(*kc.sortCond, valueSub, nameSub)
		if err != nil {
			return nil, err
		}
		kc.sortCond = k
	}
	return kc, nil
}

func parseKeyCondExprStr(sp *string) (*keyCondExpr, error) {
	if sp == nil {
		return nil, errs.Errorf("%v: %v", ErrNil, ErrInvalidKeyCondition)
	}
	s := *sp
	conds := strings.SplitN(s, " AND ", 2)
	partitionCond, err := parseKeyCond(conds[0])
	if err != nil {
		return nil, err
	}
	if partitionCond.op != eq {
		return nil, errs.Errorf("%v: partition key condition: want '=', got '%v'", ErrInvalidKeyCondition, partitionCond.op)
	}
	keyCond := &keyCondExpr{partitionCond: *partitionCond}
	if len(conds) > 1 {
		kc, err := parseKeyCond(conds[1])
		if err != nil {
			return nil, err
		}
		keyCond.sortCond = kc
	}
	return keyCond, nil
}

func parseKeyCond(s string) (*keyCond, error) {
	if t := strings.SplitN(s, "<=", 2); len(t) == 2 {
		return newKeyCond(t, lessEq)
	} else if t := strings.SplitN(s, ">=", 2); len(t) == 2 {
		return newKeyCond(t, greaterEq)
	} else if t := strings.SplitN(s, "=", 2); len(t) == 2 {
		return newKeyCond(t, eq)
	} else if t := strings.SplitN(s, "<", 2); len(t) == 2 {
		return newKeyCond(t, less)
	} else if t := strings.SplitN(s, ">", 2); len(t) == 2 {
		return newKeyCond(t, greater)
	} else if strings.HasPrefix(s, "begins_with") {
		s = strings.TrimSpace(strings.TrimPrefix(s, "begins_with"))
		if !strings.HasPrefix(s, "(") || !strings.HasSuffix(s, ")") {
			return nil, errs.Errorf("%v: invalid begins_with syntax '%s'", ErrInvalidKeyCondition, s)
		}
		s = strings.Trim(s, "()")
		return newKeyCond(strings.SplitN(s, ",", 2), beginsWith)
	} else if t := strings.SplitN(s, " BETWEEN ", 2); len(t) == 2 {
		t2 := strings.Split(t[1], " AND ")
		return newKeyCond(append([]string{t[0]}, t2...), between)
	}
	return nil, errs.Errorf("%v: %s", ErrInvalidKeyCondition, s)
}

func newKeyCond(s []string, op op) (*keyCond, error) {
	if (op != between && len(s) != 2) || (op == between && len(s) != 3) {
		return nil, errs.Errorf("%v: invalid '%v' expression", ErrInvalidKeyCondition, op)
	}
	keyName := strings.TrimSpace(s[0])
	if !reExprName.MatchString(keyName) {
		return nil, errs.Errorf("%v: invalid expression attribute name '%s' ", ErrInvalidKeyCondition, keyName)
	}
	val := strings.TrimSpace(s[1])
	if !reExprVal.MatchString(val) {
		return nil, errs.Errorf("%v: invalid expression attribute value '%s' ", ErrInvalidKeyCondition, val)
	}
	if op != between {
		return &keyCond{keyName: keyName, op: op, val: val}, nil
	}
	val2 := strings.TrimSpace(s[2])
	if !reExprVal.MatchString(val2) {
		return nil, errs.Errorf("%v: invalid expression attribute value '%s' ", ErrInvalidKeyCondition, val2)
	}
	return &keyCond{keyName: keyName, op: op, val: val, val2: val2}, nil
}

func substituteName(s string, nameSub map[string]*string) (string, error) {
	if !strings.HasPrefix(s, "#") {
		return s, nil
	}
	if nameSub == nil {
		return "", errs.Errorf("%v: %s: ExpressionAttributeNames are nil", ErrSubstitution, s)
	}
	result, ok := nameSub[s]
	if !ok {
		return "", errs.Errorf("%v: %s", ErrSubstitution, s)
	}
	return *result, nil
}

func substituteKeyCondExpr(k keyCond, valueSub Item, nameSub map[string]*string) (*keyCond, error) {
	var err error
	k.keyName, err = substituteName(k.keyName, nameSub)
	if err != nil {
		return nil, err
	}
	av, ok := valueSub[k.val]
	if !ok {
		return nil, errs.Errorf("%v: %v: %s", ErrInvalidKeyCondition, ErrSubstitution, k.val)
	}
	k.av = av
	if k.val2 != "" {
		k.av2, ok = valueSub[k.val2]
		if !ok {
			return nil, errs.Errorf("%v: %v: %s", ErrInvalidKeyCondition, ErrSubstitution, k.val2)
		}
	}
	return &k, nil
}
