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

// CasbinRuleDelete is the builder for deleting a CasbinRule entity.
type CasbinRuleDelete struct {
	config
	hooks    []Hook
	mutation *CasbinRuleMutation
}

// Where adds a new predicate to the CasbinRuleDelete builder.
func (crd *CasbinRuleDelete) Where(ps ...predicate.CasbinRule) *CasbinRuleDelete {
	crd.mutation.predicates = append(crd.mutation.predicates, ps...)
	return crd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (crd *CasbinRuleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(crd.hooks) == 0 {
		affected, err = crd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CasbinRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			crd.mutation = mutation
			affected, err = crd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(crd.hooks) - 1; i >= 0; i-- {
			mut = crd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, crd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (crd *CasbinRuleDelete) ExecX(ctx context.Context) int {
	n, err := crd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (crd *CasbinRuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: casbinrule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: casbinrule.FieldID,
			},
		},
	}
	if ps := crd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, crd.driver, _spec)
}

// CasbinRuleDeleteOne is the builder for deleting a single CasbinRule entity.
type CasbinRuleDeleteOne struct {
	crd *CasbinRuleDelete
}

// Exec executes the deletion query.
func (crdo *CasbinRuleDeleteOne) Exec(ctx context.Context) error {
	n, err := crdo.crd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{casbinrule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (crdo *CasbinRuleDeleteOne) ExecX(ctx context.Context) {
	crdo.crd.ExecX(ctx)
}
