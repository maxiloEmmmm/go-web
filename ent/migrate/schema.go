// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CasbinRulesColumns holds the columns for the "casbin_rules" table.
	CasbinRulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ptype", Type: field.TypeString, Size: 1},
		{Name: "v0", Type: field.TypeString, Nullable: true, Size: 60, Default: ""},
		{Name: "v1", Type: field.TypeString, Nullable: true, Size: 60, Default: ""},
		{Name: "v2", Type: field.TypeString, Nullable: true, Size: 60, Default: ""},
		{Name: "v3", Type: field.TypeString, Nullable: true, Size: 60, Default: ""},
		{Name: "v4", Type: field.TypeString, Nullable: true, Size: 60, Default: ""},
		{Name: "v5", Type: field.TypeString, Nullable: true, Size: 60, Default: ""},
	}
	// CasbinRulesTable holds the schema information for the "casbin_rules" table.
	CasbinRulesTable = &schema.Table{
		Name:        "casbin_rules",
		Columns:     CasbinRulesColumns,
		PrimaryKey:  []*schema.Column{CasbinRulesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CasbinRulesTable,
	}
)

func init() {
}
