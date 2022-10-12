package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Text("id").
			Unique().
			Immutable().
			NotEmpty(),
		field.Text("password").
			NotEmpty(),
		field.Text("email"),
		field.Time("create_at").
			Default(time.Now).
			Immutable(),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
