package schema

import (
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
    "github.com/facebookincubator/ent"
)
// Purpose holds the schema definition for the InstallmentEntitye entity.
type InstallmentEntity struct {
	ent.Schema
}

// Fields of the InstallmentEntity.
func (InstallmentEntitye) Fields() []ent.Field {
	return []ent.Field{
		field.String("InstallmentID").
			Unique(),
			
	}
}

// Edges of the InstallmentEntity.
func (InstallmentEntity) Edges() []ent.Edge {
	return []ent.Edge{
		
	}
}
