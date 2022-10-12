package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type Book struct {
	ent.Schema
}

func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.Text("user_id").
			NotEmpty(),
		field.Text("name").
			NotEmpty(),
		field.Float("price").
			Positive().
			Default(9999),
		field.Text("pic_url"),
		field.Time("create_at").
			Default(time.Now).
			Immutable(),
	}
}

func (Book) Edges() []ent.Edge {
	return nil
}
