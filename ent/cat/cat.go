// Code generated by ent, DO NOT EDIT.

package cat

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the cat type in the database.
	Label = "cat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeToilets holds the string denoting the toilets edge name in mutations.
	EdgeToilets = "toilets"
	// Table holds the table name of the cat in the database.
	Table = "cats"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "cats"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_cats"
	// ToiletsTable is the table that holds the toilets relation/edge.
	ToiletsTable = "toilets"
	// ToiletsInverseTable is the table name for the Toilet entity.
	// It exists in this package in order to avoid circular dependency with the "toilet" package.
	ToiletsInverseTable = "toilets"
	// ToiletsColumn is the table column denoting the toilets relation/edge.
	ToiletsColumn = "cat_toilets"
)

// Columns holds all SQL columns for cat fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldBirthday,
	FieldSex,
	FieldWeight,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cats"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_cats",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Sex defines the type for the "sex" enum field.
type Sex string

// Sex values.
const (
	SexMale   Sex = "male"
	SexFemale Sex = "female"
)

func (s Sex) String() string {
	return string(s)
}

// SexValidator is a validator for the "sex" field enum values. It is called by the builders before save.
func SexValidator(s Sex) error {
	switch s {
	case SexMale, SexFemale:
		return nil
	default:
		return fmt.Errorf("cat: invalid enum value for sex field: %q", s)
	}
}