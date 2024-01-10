package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// Metadata holds the schema definition for the Metadata entity.
type Metadata struct {
	ent.Schema
}

// Fields of the Metadata.
func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Annotations(entsql.DefaultExpr("ulid_generate()")),
		field.String("cid").Unique(),
		field.String("image"),
		field.String("description"),
		field.String("name"),
		field.Time("created_at").Default(time.Now).Annotations(entsql.DefaultExpr("CURRENT_TIMESTAMP")),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Metadata.
func (Metadata) Edges() []ent.Edge {
	return nil
}
