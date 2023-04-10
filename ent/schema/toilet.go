package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Toilet holds the schema definition for the Toilet entity.
type Toilet struct {
	ent.Schema
}

// Fields of the Toilet.
func (Toilet) Fields() []ent.Field {
	return []ent.Field{
		// IDはAutoIncrementで自動生成するため、宣言しない
		field.Time("time").Comment("トイレに行った時間"),
		field.Enum("type").Values("pee", "poo").Comment("トイレに行った種類。peeかpooのみ、peeは小便、pooは大便"),
		field.Enum("condition").Values("normal", "accident").Comment("排泄物の状態。normalかaccidentのみ、normalは通常、accidentは異常"),
		field.String("memo").Optional().Comment("メモ。nullable"),
	}
}

// Edges of the Toilet.
func (Toilet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("cat", Cat.Type).Ref("toilets").Unique().Comment("トイレをした猫。Catと1対多の関係"),
	}
}
