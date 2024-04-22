package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type User struct {
	ent.Schema
}

// Fields of the Car.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email").Unique(),
		field.String("password"),
		field.Time("create_at").Default(time.Now()),
	}
}

// Edges of the Car.
func (User) Edges() []ent.Edge {
	return nil
}
