package schema

import (
		"github.com/facebookincubator/ent"
		"github.com/facebookincubator/ent/schema/edge"
		"github.com/facebookincubator/ent/schema/field"
)
// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("idemployee").MaxLen(6).MinLen(6),

	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
	edge.From("gender",Gender.Type).Ref("users").Unique(),
	}
}
