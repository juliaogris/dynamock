package dynamock

import (
	"errors"
	"fmt"

	"foxygo.at/s/errs"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	ErrUnknownTable     = errors.New("unknown table")
	ErrMissingName      = errors.New("missing name")
	ErrSchemaValidation = errors.New("invalid schema")
	ErrUnknownType      = errs.Errorf("unknown type")
	ErrInvalidKey       = errs.Errorf("invalid key")
	ErrNil              = errs.Errorf("unexpected nil")

	ErrItemValidation   = errors.New("invalid item")
	ErrPrimaryKeyVal    = errs.Errorf("bad primary key value")
	ErrGSIVal           = errs.Errorf("bad GSI value")
	ErrMissingType      = errors.New("missing type")
	ErrInvalidType      = errors.New("invalid type")
	ErrMissingAttribute = errors.New("missing attribute")
)

func validateDB(db *DB) error {
	for _, t := range db.RawTables {
		if err := validateTable(t); err != nil {
			return err
		}
	}
	return nil
}

func validateTable(t *Table) error {
	if t.Name == "" {
		return errs.Errorf("validateTable: %v: table.Name", ErrMissingName)
	}
	if err := validateKeyDef(t.Schema.PrimaryKey); err != nil {
		return errs.Errorf("%v: %v (primary key)", ErrSchemaValidation, err)
	}
	for _, gsi := range t.Schema.GSIs {
		if gsi.Name == "" {
			return errs.Errorf("validateTable: %v: globalSecondaryIndex.Name in table %s", ErrMissingName, t.Name)
		}
		if err := validateKeyDef(gsi); err != nil {
			return errs.Errorf("%v: %v (globalSecondaryIndex %s)", ErrSchemaValidation, err, gsi.Name)
		}
	}
	for _, item := range t.items {
		if err := validateItem(item, t.Schema); err != nil {
			return errs.Errorf("validateTable: %v (table: '%s')", err, t.Name)
		}
	}
	return nil
}

func validateKeyDef(k KeyDef) error {
	if err := validateKeyPartDef(&k.PartitionKey); err != nil {
		return err
	}
	if k.SortKey == nil {
		return nil
	}
	return validateKeyPartDef(k.SortKey)
}

func validateKeyPartDef(k *KeyPartDef) error {
	if k.Name == "" {
		return ErrMissingName
	}
	return validateKeyType(k.Type)
}

func validateItem(item Item, schema Schema) error {
	if err := validateKey(item, schema.PrimaryKey); err != nil {
		return errs.New(ErrPrimaryKeyVal, err)
	}
	for _, gsi := range schema.GSIs {
		if !hasKey(item, gsi) {
			continue
		}
		if err := validateKey(item, gsi); err != nil {
			return errs.New(ErrGSIVal, err)
		}
	}
	return nil
}

func hasKey(item Item, k KeyDef) bool {
	pk := k.PartitionKey
	if item[pk.Name] == nil {
		return false
	}
	sk := k.SortKey
	if sk != nil && item[sk.Name] == nil {
		return false
	}
	return true
}

func validateKey(item Item, k KeyDef) error {
	pk := k.PartitionKey
	if err := validateAttrKeyType(item[pk.Name], pk.Type); err != nil {
		return errs.Errorf("validateKey: bad key type for '%s': %v", pk.Name, err)
	}
	sk := k.SortKey
	if sk == nil {
		return nil
	}
	return validateAttrKeyType(item[sk.Name], sk.Type)
}

func validateKeyType(typeStr string) error {
	switch typeStr {
	case "string", "number":
		return nil
	}
	return errs.Errorf("%V: validateKeyType: %s", ErrUnknownType, typeStr)
}

func validateAttrKeyType(attr *dynamodb.AttributeValue, attrType string) error {
	if attr == nil {
		return errs.Errorf("validateAttrKeyType: %v", ErrMissingAttribute)
	}
	switch {
	case attrType == "string" && attr.S != nil:
		return nil
	case attrType == "number" && attr.N != nil:
		return nil
	}
	return errs.Errorf("%v: no %s in attribute %+v", ErrMissingType, attrType, attr)
}

func validateTableName(db *DB, t *string) error {
	if t == nil {
		return errs.Errorf("%v: TableName", ErrNil)
	}
	if _, ok := db.tables[*t]; !ok {
		return fmt.Errorf("%w: %s", ErrUnknownTable, *t) // should be a dynamodb.ResourceNotFoundExcpetion
	}
	return nil
}

func validateKeyItem(key Item, schema Schema) error {
	if len(key) == 0 {
		return errs.Errorf("%v: empty key", ErrInvalidKey)
	}
	if len(key) > 2 {
		return errs.Errorf("%v: key with more than two fields", ErrInvalidKey)
	}
	return validateItem(key, schema)
}
