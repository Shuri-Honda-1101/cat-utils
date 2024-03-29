// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Shuri-Honda-1101/cat-utils/ent/cat"
	"github.com/Shuri-Honda-1101/cat-utils/ent/toilet"
	"github.com/google/uuid"
)

// Toilet is the model entity for the Toilet schema.
type Toilet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// トイレに行った時間
	Time time.Time `json:"time,omitempty"`
	// トイレに行った種類。peeかpooのみ、peeは小便、pooは大便
	Type toilet.Type `json:"type,omitempty"`
	// 備考。nullable
	Memo string `json:"memo,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ToiletQuery when eager-loading is set.
	Edges        ToiletEdges `json:"edges"`
	cat_toilets  *uuid.UUID
	selectValues sql.SelectValues
}

// ToiletEdges holds the relations/edges for other nodes in the graph.
type ToiletEdges struct {
	// トイレをした猫。Catと1対多の関係
	Cat *Cat `json:"cat,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CatOrErr returns the Cat value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ToiletEdges) CatOrErr() (*Cat, error) {
	if e.loadedTypes[0] {
		if e.Cat == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: cat.Label}
		}
		return e.Cat, nil
	}
	return nil, &NotLoadedError{edge: "cat"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Toilet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case toilet.FieldID:
			values[i] = new(sql.NullInt64)
		case toilet.FieldType, toilet.FieldMemo:
			values[i] = new(sql.NullString)
		case toilet.FieldTime:
			values[i] = new(sql.NullTime)
		case toilet.ForeignKeys[0]: // cat_toilets
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Toilet fields.
func (t *Toilet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case toilet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case toilet.FieldTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				t.Time = value.Time
			}
		case toilet.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = toilet.Type(value.String)
			}
		case toilet.FieldMemo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memo", values[i])
			} else if value.Valid {
				t.Memo = value.String
			}
		case toilet.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field cat_toilets", values[i])
			} else if value.Valid {
				t.cat_toilets = new(uuid.UUID)
				*t.cat_toilets = *value.S.(*uuid.UUID)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Toilet.
// This includes values selected through modifiers, order, etc.
func (t *Toilet) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryCat queries the "cat" edge of the Toilet entity.
func (t *Toilet) QueryCat() *CatQuery {
	return NewToiletClient(t.config).QueryCat(t)
}

// Update returns a builder for updating this Toilet.
// Note that you need to call Toilet.Unwrap() before calling this method if this Toilet
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Toilet) Update() *ToiletUpdateOne {
	return NewToiletClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Toilet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Toilet) Unwrap() *Toilet {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Toilet is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Toilet) String() string {
	var builder strings.Builder
	builder.WriteString("Toilet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("time=")
	builder.WriteString(t.Time.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", t.Type))
	builder.WriteString(", ")
	builder.WriteString("memo=")
	builder.WriteString(t.Memo)
	builder.WriteByte(')')
	return builder.String()
}

// Toilets is a parsable slice of Toilet.
type Toilets []*Toilet
