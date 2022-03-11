package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PrivateMessage holds the schema definition for the PrivateMessage entity.
type PrivateMessage struct {
	ent.Schema
}

// Fields of the PrivateMessage.
func (PrivateMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("content").NotEmpty(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the PrivateMessage.
func (PrivateMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("private_messages").
			Unique().
			Required(),
		edge.From("recipient", User.Type).
			Ref("private_messages_received").
			Unique().
			Required(),
		edge.To("children", PrivateMessage.Type).
			From("parent").
			Unique(),
	}
}
