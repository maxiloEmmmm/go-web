// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/maxiloEmmmm/go-web/ent/casbinrule"
	"github.com/maxiloEmmmm/go-web/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	casbinruleFields := schema.CasbinRule{}.Fields()
	_ = casbinruleFields
	// casbinruleDescPType is the schema descriptor for PType field.
	casbinruleDescPType := casbinruleFields[0].Descriptor()
	// casbinrule.PTypeValidator is a validator for the "PType" field. It is called by the builders before save.
	casbinrule.PTypeValidator = casbinruleDescPType.Validators[0].(func(string) error)
	// casbinruleDescV0 is the schema descriptor for v0 field.
	casbinruleDescV0 := casbinruleFields[1].Descriptor()
	// casbinrule.DefaultV0 holds the default value on creation for the v0 field.
	casbinrule.DefaultV0 = casbinruleDescV0.Default.(string)
	// casbinrule.V0Validator is a validator for the "v0" field. It is called by the builders before save.
	casbinrule.V0Validator = casbinruleDescV0.Validators[0].(func(string) error)
	// casbinruleDescV1 is the schema descriptor for v1 field.
	casbinruleDescV1 := casbinruleFields[2].Descriptor()
	// casbinrule.DefaultV1 holds the default value on creation for the v1 field.
	casbinrule.DefaultV1 = casbinruleDescV1.Default.(string)
	// casbinrule.V1Validator is a validator for the "v1" field. It is called by the builders before save.
	casbinrule.V1Validator = casbinruleDescV1.Validators[0].(func(string) error)
	// casbinruleDescV2 is the schema descriptor for v2 field.
	casbinruleDescV2 := casbinruleFields[3].Descriptor()
	// casbinrule.DefaultV2 holds the default value on creation for the v2 field.
	casbinrule.DefaultV2 = casbinruleDescV2.Default.(string)
	// casbinrule.V2Validator is a validator for the "v2" field. It is called by the builders before save.
	casbinrule.V2Validator = casbinruleDescV2.Validators[0].(func(string) error)
	// casbinruleDescV3 is the schema descriptor for v3 field.
	casbinruleDescV3 := casbinruleFields[4].Descriptor()
	// casbinrule.DefaultV3 holds the default value on creation for the v3 field.
	casbinrule.DefaultV3 = casbinruleDescV3.Default.(string)
	// casbinrule.V3Validator is a validator for the "v3" field. It is called by the builders before save.
	casbinrule.V3Validator = casbinruleDescV3.Validators[0].(func(string) error)
	// casbinruleDescV4 is the schema descriptor for v4 field.
	casbinruleDescV4 := casbinruleFields[5].Descriptor()
	// casbinrule.DefaultV4 holds the default value on creation for the v4 field.
	casbinrule.DefaultV4 = casbinruleDescV4.Default.(string)
	// casbinrule.V4Validator is a validator for the "v4" field. It is called by the builders before save.
	casbinrule.V4Validator = casbinruleDescV4.Validators[0].(func(string) error)
	// casbinruleDescV5 is the schema descriptor for v5 field.
	casbinruleDescV5 := casbinruleFields[6].Descriptor()
	// casbinrule.DefaultV5 holds the default value on creation for the v5 field.
	casbinrule.DefaultV5 = casbinruleDescV5.Default.(string)
	// casbinrule.V5Validator is a validator for the "v5" field. It is called by the builders before save.
	casbinrule.V5Validator = casbinruleDescV5.Validators[0].(func(string) error)
}
