// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/maxiloEmmmm/go-web/ent/casbinrule"
	"github.com/maxiloEmmmm/go-web/ent/predicate"
)

// CasbinRuleUpdate is the builder for updating CasbinRule entities.
type CasbinRuleUpdate struct {
	config
	hooks    []Hook
	mutation *CasbinRuleMutation
}

// Where adds a new predicate for the CasbinRuleUpdate builder.
func (cru *CasbinRuleUpdate) Where(ps ...predicate.CasbinRule) *CasbinRuleUpdate {
	cru.mutation.predicates = append(cru.mutation.predicates, ps...)
	return cru
}

// SetPType sets the "PType" field.
func (cru *CasbinRuleUpdate) SetPType(s string) *CasbinRuleUpdate {
	cru.mutation.SetPType(s)
	return cru
}

// SetV0 sets the "v0" field.
func (cru *CasbinRuleUpdate) SetV0(s string) *CasbinRuleUpdate {
	cru.mutation.SetV0(s)
	return cru
}

// SetNillableV0 sets the "v0" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV0(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV0(*s)
	}
	return cru
}

// ClearV0 clears the value of the "v0" field.
func (cru *CasbinRuleUpdate) ClearV0() *CasbinRuleUpdate {
	cru.mutation.ClearV0()
	return cru
}

// SetV1 sets the "v1" field.
func (cru *CasbinRuleUpdate) SetV1(s string) *CasbinRuleUpdate {
	cru.mutation.SetV1(s)
	return cru
}

// SetNillableV1 sets the "v1" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV1(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV1(*s)
	}
	return cru
}

// ClearV1 clears the value of the "v1" field.
func (cru *CasbinRuleUpdate) ClearV1() *CasbinRuleUpdate {
	cru.mutation.ClearV1()
	return cru
}

// SetV2 sets the "v2" field.
func (cru *CasbinRuleUpdate) SetV2(s string) *CasbinRuleUpdate {
	cru.mutation.SetV2(s)
	return cru
}

// SetNillableV2 sets the "v2" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV2(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV2(*s)
	}
	return cru
}

// ClearV2 clears the value of the "v2" field.
func (cru *CasbinRuleUpdate) ClearV2() *CasbinRuleUpdate {
	cru.mutation.ClearV2()
	return cru
}

// SetV3 sets the "v3" field.
func (cru *CasbinRuleUpdate) SetV3(s string) *CasbinRuleUpdate {
	cru.mutation.SetV3(s)
	return cru
}

// SetNillableV3 sets the "v3" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV3(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV3(*s)
	}
	return cru
}

// ClearV3 clears the value of the "v3" field.
func (cru *CasbinRuleUpdate) ClearV3() *CasbinRuleUpdate {
	cru.mutation.ClearV3()
	return cru
}

// SetV4 sets the "v4" field.
func (cru *CasbinRuleUpdate) SetV4(s string) *CasbinRuleUpdate {
	cru.mutation.SetV4(s)
	return cru
}

// SetNillableV4 sets the "v4" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV4(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV4(*s)
	}
	return cru
}

// ClearV4 clears the value of the "v4" field.
func (cru *CasbinRuleUpdate) ClearV4() *CasbinRuleUpdate {
	cru.mutation.ClearV4()
	return cru
}

// SetV5 sets the "v5" field.
func (cru *CasbinRuleUpdate) SetV5(s string) *CasbinRuleUpdate {
	cru.mutation.SetV5(s)
	return cru
}

// SetNillableV5 sets the "v5" field if the given value is not nil.
func (cru *CasbinRuleUpdate) SetNillableV5(s *string) *CasbinRuleUpdate {
	if s != nil {
		cru.SetV5(*s)
	}
	return cru
}

// ClearV5 clears the value of the "v5" field.
func (cru *CasbinRuleUpdate) ClearV5() *CasbinRuleUpdate {
	cru.mutation.ClearV5()
	return cru
}

// Mutation returns the CasbinRuleMutation object of the builder.
func (cru *CasbinRuleUpdate) Mutation() *CasbinRuleMutation {
	return cru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cru *CasbinRuleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cru.hooks) == 0 {
		if err = cru.check(); err != nil {
			return 0, err
		}
		affected, err = cru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CasbinRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cru.check(); err != nil {
				return 0, err
			}
			cru.mutation = mutation
			affected, err = cru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cru.hooks) - 1; i >= 0; i-- {
			mut = cru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cru *CasbinRuleUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *CasbinRuleUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *CasbinRuleUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cru *CasbinRuleUpdate) check() error {
	if v, ok := cru.mutation.PType(); ok {
		if err := casbinrule.PTypeValidator(v); err != nil {
			return &ValidationError{Name: "PType", err: fmt.Errorf("ent: validator failed for field \"PType\": %w", err)}
		}
	}
	if v, ok := cru.mutation.V0(); ok {
		if err := casbinrule.V0Validator(v); err != nil {
			return &ValidationError{Name: "v0", err: fmt.Errorf("ent: validator failed for field \"v0\": %w", err)}
		}
	}
	if v, ok := cru.mutation.V1(); ok {
		if err := casbinrule.V1Validator(v); err != nil {
			return &ValidationError{Name: "v1", err: fmt.Errorf("ent: validator failed for field \"v1\": %w", err)}
		}
	}
	if v, ok := cru.mutation.V2(); ok {
		if err := casbinrule.V2Validator(v); err != nil {
			return &ValidationError{Name: "v2", err: fmt.Errorf("ent: validator failed for field \"v2\": %w", err)}
		}
	}
	if v, ok := cru.mutation.V3(); ok {
		if err := casbinrule.V3Validator(v); err != nil {
			return &ValidationError{Name: "v3", err: fmt.Errorf("ent: validator failed for field \"v3\": %w", err)}
		}
	}
	if v, ok := cru.mutation.V4(); ok {
		if err := casbinrule.V4Validator(v); err != nil {
			return &ValidationError{Name: "v4", err: fmt.Errorf("ent: validator failed for field \"v4\": %w", err)}
		}
	}
	if v, ok := cru.mutation.V5(); ok {
		if err := casbinrule.V5Validator(v); err != nil {
			return &ValidationError{Name: "v5", err: fmt.Errorf("ent: validator failed for field \"v5\": %w", err)}
		}
	}
	return nil
}

func (cru *CasbinRuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   casbinrule.Table,
			Columns: casbinrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: casbinrule.FieldID,
			},
		},
	}
	if ps := cru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cru.mutation.PType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldPType,
		})
	}
	if value, ok := cru.mutation.V0(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV0,
		})
	}
	if cru.mutation.V0Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV0,
		})
	}
	if value, ok := cru.mutation.V1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV1,
		})
	}
	if cru.mutation.V1Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV1,
		})
	}
	if value, ok := cru.mutation.V2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV2,
		})
	}
	if cru.mutation.V2Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV2,
		})
	}
	if value, ok := cru.mutation.V3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV3,
		})
	}
	if cru.mutation.V3Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV3,
		})
	}
	if value, ok := cru.mutation.V4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV4,
		})
	}
	if cru.mutation.V4Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV4,
		})
	}
	if value, ok := cru.mutation.V5(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV5,
		})
	}
	if cru.mutation.V5Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV5,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{casbinrule.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CasbinRuleUpdateOne is the builder for updating a single CasbinRule entity.
type CasbinRuleUpdateOne struct {
	config
	hooks    []Hook
	mutation *CasbinRuleMutation
}

// SetPType sets the "PType" field.
func (cruo *CasbinRuleUpdateOne) SetPType(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetPType(s)
	return cruo
}

// SetV0 sets the "v0" field.
func (cruo *CasbinRuleUpdateOne) SetV0(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV0(s)
	return cruo
}

// SetNillableV0 sets the "v0" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV0(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV0(*s)
	}
	return cruo
}

// ClearV0 clears the value of the "v0" field.
func (cruo *CasbinRuleUpdateOne) ClearV0() *CasbinRuleUpdateOne {
	cruo.mutation.ClearV0()
	return cruo
}

// SetV1 sets the "v1" field.
func (cruo *CasbinRuleUpdateOne) SetV1(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV1(s)
	return cruo
}

// SetNillableV1 sets the "v1" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV1(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV1(*s)
	}
	return cruo
}

// ClearV1 clears the value of the "v1" field.
func (cruo *CasbinRuleUpdateOne) ClearV1() *CasbinRuleUpdateOne {
	cruo.mutation.ClearV1()
	return cruo
}

// SetV2 sets the "v2" field.
func (cruo *CasbinRuleUpdateOne) SetV2(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV2(s)
	return cruo
}

// SetNillableV2 sets the "v2" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV2(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV2(*s)
	}
	return cruo
}

// ClearV2 clears the value of the "v2" field.
func (cruo *CasbinRuleUpdateOne) ClearV2() *CasbinRuleUpdateOne {
	cruo.mutation.ClearV2()
	return cruo
}

// SetV3 sets the "v3" field.
func (cruo *CasbinRuleUpdateOne) SetV3(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV3(s)
	return cruo
}

// SetNillableV3 sets the "v3" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV3(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV3(*s)
	}
	return cruo
}

// ClearV3 clears the value of the "v3" field.
func (cruo *CasbinRuleUpdateOne) ClearV3() *CasbinRuleUpdateOne {
	cruo.mutation.ClearV3()
	return cruo
}

// SetV4 sets the "v4" field.
func (cruo *CasbinRuleUpdateOne) SetV4(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV4(s)
	return cruo
}

// SetNillableV4 sets the "v4" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV4(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV4(*s)
	}
	return cruo
}

// ClearV4 clears the value of the "v4" field.
func (cruo *CasbinRuleUpdateOne) ClearV4() *CasbinRuleUpdateOne {
	cruo.mutation.ClearV4()
	return cruo
}

// SetV5 sets the "v5" field.
func (cruo *CasbinRuleUpdateOne) SetV5(s string) *CasbinRuleUpdateOne {
	cruo.mutation.SetV5(s)
	return cruo
}

// SetNillableV5 sets the "v5" field if the given value is not nil.
func (cruo *CasbinRuleUpdateOne) SetNillableV5(s *string) *CasbinRuleUpdateOne {
	if s != nil {
		cruo.SetV5(*s)
	}
	return cruo
}

// ClearV5 clears the value of the "v5" field.
func (cruo *CasbinRuleUpdateOne) ClearV5() *CasbinRuleUpdateOne {
	cruo.mutation.ClearV5()
	return cruo
}

// Mutation returns the CasbinRuleMutation object of the builder.
func (cruo *CasbinRuleUpdateOne) Mutation() *CasbinRuleMutation {
	return cruo.mutation
}

// Save executes the query and returns the updated CasbinRule entity.
func (cruo *CasbinRuleUpdateOne) Save(ctx context.Context) (*CasbinRule, error) {
	var (
		err  error
		node *CasbinRule
	)
	if len(cruo.hooks) == 0 {
		if err = cruo.check(); err != nil {
			return nil, err
		}
		node, err = cruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CasbinRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cruo.check(); err != nil {
				return nil, err
			}
			cruo.mutation = mutation
			node, err = cruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cruo.hooks) - 1; i >= 0; i-- {
			mut = cruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *CasbinRuleUpdateOne) SaveX(ctx context.Context) *CasbinRule {
	node, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cruo *CasbinRuleUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *CasbinRuleUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cruo *CasbinRuleUpdateOne) check() error {
	if v, ok := cruo.mutation.PType(); ok {
		if err := casbinrule.PTypeValidator(v); err != nil {
			return &ValidationError{Name: "PType", err: fmt.Errorf("ent: validator failed for field \"PType\": %w", err)}
		}
	}
	if v, ok := cruo.mutation.V0(); ok {
		if err := casbinrule.V0Validator(v); err != nil {
			return &ValidationError{Name: "v0", err: fmt.Errorf("ent: validator failed for field \"v0\": %w", err)}
		}
	}
	if v, ok := cruo.mutation.V1(); ok {
		if err := casbinrule.V1Validator(v); err != nil {
			return &ValidationError{Name: "v1", err: fmt.Errorf("ent: validator failed for field \"v1\": %w", err)}
		}
	}
	if v, ok := cruo.mutation.V2(); ok {
		if err := casbinrule.V2Validator(v); err != nil {
			return &ValidationError{Name: "v2", err: fmt.Errorf("ent: validator failed for field \"v2\": %w", err)}
		}
	}
	if v, ok := cruo.mutation.V3(); ok {
		if err := casbinrule.V3Validator(v); err != nil {
			return &ValidationError{Name: "v3", err: fmt.Errorf("ent: validator failed for field \"v3\": %w", err)}
		}
	}
	if v, ok := cruo.mutation.V4(); ok {
		if err := casbinrule.V4Validator(v); err != nil {
			return &ValidationError{Name: "v4", err: fmt.Errorf("ent: validator failed for field \"v4\": %w", err)}
		}
	}
	if v, ok := cruo.mutation.V5(); ok {
		if err := casbinrule.V5Validator(v); err != nil {
			return &ValidationError{Name: "v5", err: fmt.Errorf("ent: validator failed for field \"v5\": %w", err)}
		}
	}
	return nil
}

func (cruo *CasbinRuleUpdateOne) sqlSave(ctx context.Context) (_node *CasbinRule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   casbinrule.Table,
			Columns: casbinrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: casbinrule.FieldID,
			},
		},
	}
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing CasbinRule.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := cruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cruo.mutation.PType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldPType,
		})
	}
	if value, ok := cruo.mutation.V0(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV0,
		})
	}
	if cruo.mutation.V0Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV0,
		})
	}
	if value, ok := cruo.mutation.V1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV1,
		})
	}
	if cruo.mutation.V1Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV1,
		})
	}
	if value, ok := cruo.mutation.V2(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV2,
		})
	}
	if cruo.mutation.V2Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV2,
		})
	}
	if value, ok := cruo.mutation.V3(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV3,
		})
	}
	if cruo.mutation.V3Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV3,
		})
	}
	if value, ok := cruo.mutation.V4(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV4,
		})
	}
	if cruo.mutation.V4Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV4,
		})
	}
	if value, ok := cruo.mutation.V5(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: casbinrule.FieldV5,
		})
	}
	if cruo.mutation.V5Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: casbinrule.FieldV5,
		})
	}
	_node = &CasbinRule{config: cruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{casbinrule.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
