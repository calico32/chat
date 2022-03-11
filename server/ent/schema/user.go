package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Immutable(),
		field.String("username").NotEmpty().Immutable(),
		field.String("avatar_url").Optional().Nillable(),
		field.String("secret_hash").Optional().Nillable().Sensitive().Immutable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
		edge.To("private_messages", PrivateMessage.Type),
		edge.To("private_messages_received", PrivateMessage.Type),
	}
}
