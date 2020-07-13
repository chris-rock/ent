// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/galaxy"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/planet"
	"github.com/facebookincubator/ent/schema/field"
)

// GalaxyCreate is the builder for creating a Galaxy entity.
type GalaxyCreate struct {
	config
	mutation *GalaxyMutation
	hooks    []Hook
}

// SetName sets the name field.
func (gc *GalaxyCreate) SetName(s string) *GalaxyCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetType sets the type field.
func (gc *GalaxyCreate) SetType(ga galaxy.Type) *GalaxyCreate {
	gc.mutation.SetType(ga)
	return gc
}

// AddPlanetIDs adds the planets edge to Planet by ids.
func (gc *GalaxyCreate) AddPlanetIDs(ids ...int) *GalaxyCreate {
	gc.mutation.AddPlanetIDs(ids...)
	return gc
}

// AddPlanets adds the planets edges to Planet.
func (gc *GalaxyCreate) AddPlanets(p ...*Planet) *GalaxyCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return gc.AddPlanetIDs(ids...)
}

// Mutation returns the GalaxyMutation object of the builder.
func (gc *GalaxyCreate) Mutation() *GalaxyMutation {
	return gc.mutation
}

// Save creates the Galaxy in the database.
func (gc *GalaxyCreate) Save(ctx context.Context) (*Galaxy, error) {
	if _, ok := gc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := gc.mutation.Name(); ok {
		if err := galaxy.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := gc.mutation.GetType(); !ok {
		return nil, &ValidationError{Name: "type", err: errors.New("ent: missing required field \"type\"")}
	}
	if v, ok := gc.mutation.GetType(); ok {
		if err := galaxy.TypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	var (
		err  error
		node *Galaxy
	)
	if len(gc.hooks) == 0 {
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GalaxyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gc.mutation = mutation
			node, err = gc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GalaxyCreate) SaveX(ctx context.Context) *Galaxy {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gc *GalaxyCreate) sqlSave(ctx context.Context) (*Galaxy, error) {
	ga, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ga.ID = int(id)
	return ga, nil
}

func (gc *GalaxyCreate) createSpec() (*Galaxy, *sqlgraph.CreateSpec) {
	var (
		ga    = &Galaxy{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: galaxy.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: galaxy.FieldID,
			},
		}
	)
	if value, ok := gc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: galaxy.FieldName,
		})
		ga.Name = value
	}
	if value, ok := gc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: galaxy.FieldType,
		})
		ga.Type = value
	}
	if nodes := gc.mutation.PlanetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   galaxy.PlanetsTable,
			Columns: []string{galaxy.PlanetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: planet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return ga, _spec
}
