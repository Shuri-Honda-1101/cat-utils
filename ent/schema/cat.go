package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Cat holds the schema definition for the Cat entity.
type Cat struct {
	ent.Schema
}

// Fields of the Cat.
func (Cat) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Comment("猫のID。UUIDで生成される変更不可なフィールド"),
		field.String("name").NotEmpty().MaxLen(15).Comment("猫の名前。空文字不可。最大長は15文字。"),
		field.Time("birthday").Optional().Comment("猫の誕生日、nullable"),
		field.Enum("sex").Values("male", "female").Comment("猫の性別。maleかfemaleのみ"),
		field.Int("weight").Optional().Comment("猫の体重。nullable"),
	}
}

// Edges of the Cat.
func (Cat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("cats").Unique().Comment("猫のオーナー。Userと1対多の関係"),
	}
}
