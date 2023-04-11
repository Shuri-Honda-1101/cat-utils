package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
// エンティティフィールド (またはプロパティ) 。例えば、 User の名前や年齢。
// すべてのフィールドはデフォルトでは必須であり、Optionalメソッドを使用してオプションに設定することができます。(https://entgo.io/ja/docs/schema-fields)
func (User) Fields() []ent.Field {
	return []ent.Field{
		// idフィールドはスキーマに組み込まれているため、宣言は不要です。SQLベースのデータベースでは、その型はデフォルトでint（ただしcodegenオプションで変更可能）であり、データベース内で自動インクリメントされる。(https://entgo.io/ja/docs/code-gen/#%E3%82%B3%E3%83%BC%E3%83%89%E7%94%9F%E6%88%90%E3%82%AA%E3%83%97%E3%82%B7%E3%83%A7%E3%83%B3)
		// idフィールドが全テーブルで一意になるように設定するには、スキーマ移行を実行する際にWithGlobalUniqueIDオプションを使用します。(https://entgo.io/ja/docs/code-gen/#%E3%82%B3%E3%83%BC%E3%83%89%E7%94%9F%E6%88%90%E3%82%AA%E3%83%97%E3%82%B7%E3%83%A7%E3%83%B3)
		// id フィールドに異なる設定が必要な場合、またはアプリケーションによるエンティティ作成時に id 値を提供する必要がある場合（UUID など）、組み込みの id 設定をオーバーライドします。
		// IDを生成するためのカスタム関数を設定する必要がある場合は、DefaultFuncを使用して、リソースの作成時に常に実行される関数を指定することができます(https://entgo.io/ja/docs/faq/#id%E3%81%AE%E3%82%AB%E3%82%B9%E3%82%BF%E3%83%A0%E3%82%B8%E3%82%A7%E3%83%8D%E3%83%AC%E3%83%BC%E3%82%BF%E3%82%92%E4%BD%BF%E7%94%A8%E3%81%99%E3%82%8B%E6%96%B9%E6%B3%95%E3%81%AF)
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Comment("ユーザーID。UUIDで生成される変更不可なフィールド"),
		field.String("name").NotEmpty().MaxLen(15).Comment("ユーザー名。空文字不可。最大長は15文字。"),
		field.String("email").NotEmpty().Unique().Comment("メールアドレス。空文字不可。一意。"),
		field.Time("created_at").
			Default(time.Now).Immutable().Comment("作成日時。デフォルトは現在時刻。変更不可。"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).Comment("更新日時。デフォルトは現在時刻。更新時は現在時刻。"),
		field.String("password").Sensitive().Comment("パスワード。センシティブ"),
	}
}

// Edges of the User.
// エンティティエッジ（またはリレーション）。例えば、 Userのグループ、 Userのフレンド。
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cats", Cat.Type),
	}
}

// データベース固有のオプション。例えば、インデックスや一意インデックスなど。
func (User) Indexes() []ent.Index {
	return nil
}
