package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// CasbinRule holds the schema definition for the CasbinRule entity.
type CasbinRule struct {
	ent.Schema
}

// Fields of the CasbinRule.
func (CasbinRule) Fields() []ent.Field {
	return []ent.Field{
		field.String("PType").MaxLen(40),
		field.String("v0").MaxLen(40).Default("").Optional(),
		field.String("v1").MaxLen(40).Default("").Optional(),
		field.String("v2").MaxLen(40).Default("").Optional(),
		field.String("v3").MaxLen(40).Default("").Optional(),
		field.String("v4").MaxLen(40).Default("").Optional(),
		field.String("v5").MaxLen(40).Default("").Optional(),
	}
}