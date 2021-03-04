// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/maxiloEmmmm/go-web/ent/casbinrule"
)

// CasbinRule is the model entity for the CasbinRule schema.
type CasbinRule struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PType holds the value of the "PType" field.
	PType string `json:"PType,omitempty"`
	// V0 holds the value of the "v0" field.
	V0 string `json:"v0,omitempty"`
	// V1 holds the value of the "v1" field.
	V1 string `json:"v1,omitempty"`
	// V2 holds the value of the "v2" field.
	V2 string `json:"v2,omitempty"`
	// V3 holds the value of the "v3" field.
	V3 string `json:"v3,omitempty"`
	// V4 holds the value of the "v4" field.
	V4 string `json:"v4,omitempty"`
	// V5 holds the value of the "v5" field.
	V5 string `json:"v5,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CasbinRule) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case casbinrule.FieldID:
			values[i] = &sql.NullInt64{}
		case casbinrule.FieldPType, casbinrule.FieldV0, casbinrule.FieldV1, casbinrule.FieldV2, casbinrule.FieldV3, casbinrule.FieldV4, casbinrule.FieldV5:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type CasbinRule", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CasbinRule fields.
func (cr *CasbinRule) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case casbinrule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cr.ID = int(value.Int64)
		case casbinrule.FieldPType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PType", values[i])
			} else if value.Valid {
				cr.PType = value.String
			}
		case casbinrule.FieldV0:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field v0", values[i])
			} else if value.Valid {
				cr.V0 = value.String
			}
		case casbinrule.FieldV1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field v1", values[i])
			} else if value.Valid {
				cr.V1 = value.String
			}
		case casbinrule.FieldV2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field v2", values[i])
			} else if value.Valid {
				cr.V2 = value.String
			}
		case casbinrule.FieldV3:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field v3", values[i])
			} else if value.Valid {
				cr.V3 = value.String
			}
		case casbinrule.FieldV4:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field v4", values[i])
			} else if value.Valid {
				cr.V4 = value.String
			}
		case casbinrule.FieldV5:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field v5", values[i])
			} else if value.Valid {
				cr.V5 = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CasbinRule.
// Note that you need to call CasbinRule.Unwrap() before calling this method if this CasbinRule
// was returned from a transaction, and the transaction was committed or rolled back.
func (cr *CasbinRule) Update() *CasbinRuleUpdateOne {
	return (&CasbinRuleClient{config: cr.config}).UpdateOne(cr)
}

// Unwrap unwraps the CasbinRule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cr *CasbinRule) Unwrap() *CasbinRule {
	tx, ok := cr.config.driver.(*txDriver)
	if !ok {
		panic("ent: CasbinRule is not a transactional entity")
	}
	cr.config.driver = tx.drv
	return cr
}

// String implements the fmt.Stringer.
func (cr *CasbinRule) String() string {
	var builder strings.Builder
	builder.WriteString("CasbinRule(")
	builder.WriteString(fmt.Sprintf("id=%v", cr.ID))
	builder.WriteString(", PType=")
	builder.WriteString(cr.PType)
	builder.WriteString(", v0=")
	builder.WriteString(cr.V0)
	builder.WriteString(", v1=")
	builder.WriteString(cr.V1)
	builder.WriteString(", v2=")
	builder.WriteString(cr.V2)
	builder.WriteString(", v3=")
	builder.WriteString(cr.V3)
	builder.WriteString(", v4=")
	builder.WriteString(cr.V4)
	builder.WriteString(", v5=")
	builder.WriteString(cr.V5)
	builder.WriteByte(')')
	return builder.String()
}

// CasbinRules is a parsable slice of CasbinRule.
type CasbinRules []*CasbinRule

func (cr CasbinRules) config(cfg config) {
	for _i := range cr {
		cr[_i].config = cfg
	}
}
