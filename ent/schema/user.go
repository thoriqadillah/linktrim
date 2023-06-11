package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name"),
		field.String("email"),
		field.String("password"),
		field.String("profile_pic").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("links", Link.Type),
	}
}
