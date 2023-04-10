// Code generated by ent, DO NOT EDIT.

package cat

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Shuri-Honda-1101/cat-utils/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Cat {
	return predicate.Cat(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Cat {
	return predicate.Cat(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Cat {
	return predicate.Cat(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Cat {
	return predicate.Cat(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Cat {
	return predicate.Cat(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Cat {
	return predicate.Cat(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Cat {
	return predicate.Cat(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Cat {
	return predicate.Cat(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Cat {
	return predicate.Cat(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Cat {
	return predicate.Cat(sql.FieldEQ(FieldName, v))
}

// Birthday applies equality check predicate on the "birthday" field. It's identical to BirthdayEQ.
func Birthday(v time.Time) predicate.Cat {
	return predicate.Cat(sql.FieldEQ(FieldBirthday, v))
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v int) predicate.Cat {
	return predicate.Cat(sql.FieldEQ(FieldWeight, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Cat {
	return predicate.Cat(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Cat {
	return predicate.Cat(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Cat {
	return predicate.Cat(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Cat {
	return predicate.Cat(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Cat {
	return predicate.Cat(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Cat {
	return predicate.Cat(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Cat {
	return predicate.Cat(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Cat {
	return predicate.Cat(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Cat {
	return predicate.Cat(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Cat {
	return predicate.Cat(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Cat {
	return predicate.Cat(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Cat {
	return predicate.Cat(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Cat {
	return predicate.Cat(sql.FieldContainsFold(FieldName, v))
}

// BirthdayEQ applies the EQ predicate on the "birthday" field.
func BirthdayEQ(v time.Time) predicate.Cat {
	return predicate.Cat(sql.FieldEQ(FieldBirthday, v))
}

// BirthdayNEQ applies the NEQ predicate on the "birthday" field.
func BirthdayNEQ(v time.Time) predicate.Cat {
	return predicate.Cat(sql.FieldNEQ(FieldBirthday, v))
}

// BirthdayIn applies the In predicate on the "birthday" field.
func BirthdayIn(vs ...time.Time) predicate.Cat {
	return predicate.Cat(sql.FieldIn(FieldBirthday, vs...))
}

// BirthdayNotIn applies the NotIn predicate on the "birthday" field.
func BirthdayNotIn(vs ...time.Time) predicate.Cat {
	return predicate.Cat(sql.FieldNotIn(FieldBirthday, vs...))
}

// BirthdayGT applies the GT predicate on the "birthday" field.
func BirthdayGT(v time.Time) predicate.Cat {
	return predicate.Cat(sql.FieldGT(FieldBirthday, v))
}

// BirthdayGTE applies the GTE predicate on the "birthday" field.
func BirthdayGTE(v time.Time) predicate.Cat {
	return predicate.Cat(sql.FieldGTE(FieldBirthday, v))
}

// BirthdayLT applies the LT predicate on the "birthday" field.
func BirthdayLT(v time.Time) predicate.Cat {
	return predicate.Cat(sql.FieldLT(FieldBirthday, v))
}

// BirthdayLTE applies the LTE predicate on the "birthday" field.
func BirthdayLTE(v time.Time) predicate.Cat {
	return predicate.Cat(sql.FieldLTE(FieldBirthday, v))
}

// BirthdayIsNil applies the IsNil predicate on the "birthday" field.
func BirthdayIsNil() predicate.Cat {
	return predicate.Cat(sql.FieldIsNull(FieldBirthday))
}

// BirthdayNotNil applies the NotNil predicate on the "birthday" field.
func BirthdayNotNil() predicate.Cat {
	return predicate.Cat(sql.FieldNotNull(FieldBirthday))
}

// SexEQ applies the EQ predicate on the "sex" field.
func SexEQ(v Sex) predicate.Cat {
	return predicate.Cat(sql.FieldEQ(FieldSex, v))
}

// SexNEQ applies the NEQ predicate on the "sex" field.
func SexNEQ(v Sex) predicate.Cat {
	return predicate.Cat(sql.FieldNEQ(FieldSex, v))
}

// SexIn applies the In predicate on the "sex" field.
func SexIn(vs ...Sex) predicate.Cat {
	return predicate.Cat(sql.FieldIn(FieldSex, vs...))
}

// SexNotIn applies the NotIn predicate on the "sex" field.
func SexNotIn(vs ...Sex) predicate.Cat {
	return predicate.Cat(sql.FieldNotIn(FieldSex, vs...))
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v int) predicate.Cat {
	return predicate.Cat(sql.FieldEQ(FieldWeight, v))
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v int) predicate.Cat {
	return predicate.Cat(sql.FieldNEQ(FieldWeight, v))
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...int) predicate.Cat {
	return predicate.Cat(sql.FieldIn(FieldWeight, vs...))
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...int) predicate.Cat {
	return predicate.Cat(sql.FieldNotIn(FieldWeight, vs...))
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v int) predicate.Cat {
	return predicate.Cat(sql.FieldGT(FieldWeight, v))
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v int) predicate.Cat {
	return predicate.Cat(sql.FieldGTE(FieldWeight, v))
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v int) predicate.Cat {
	return predicate.Cat(sql.FieldLT(FieldWeight, v))
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v int) predicate.Cat {
	return predicate.Cat(sql.FieldLTE(FieldWeight, v))
}

// WeightIsNil applies the IsNil predicate on the "weight" field.
func WeightIsNil() predicate.Cat {
	return predicate.Cat(sql.FieldIsNull(FieldWeight))
}

// WeightNotNil applies the NotNil predicate on the "weight" field.
func WeightNotNil() predicate.Cat {
	return predicate.Cat(sql.FieldNotNull(FieldWeight))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Cat {
	return predicate.Cat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Cat {
	return predicate.Cat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasToilets applies the HasEdge predicate on the "toilets" edge.
func HasToilets() predicate.Cat {
	return predicate.Cat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ToiletsTable, ToiletsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasToiletsWith applies the HasEdge predicate on the "toilets" edge with a given conditions (other predicates).
func HasToiletsWith(preds ...predicate.Toilet) predicate.Cat {
	return predicate.Cat(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ToiletsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ToiletsTable, ToiletsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Cat) predicate.Cat {
	return predicate.Cat(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Cat) predicate.Cat {
	return predicate.Cat(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Cat) predicate.Cat {
	return predicate.Cat(func(s *sql.Selector) {
		p(s.Not())
	})
}