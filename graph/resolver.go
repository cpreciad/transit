package graph

import (
	"github.com/cpreciad/transit/graph/model"
	queryengine "github.com/cpreciad/transit/query_engine"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	queryEngine queryengine.QueryEngine
	// TODO: update these to store data more efficiently
	// these don't need to be updated nearly as often, so
	// a data structure that allows quicker loading will be better
	// map[id][]*model.Operator
	operators []*model.Operator
	// map[id][]*model.Line
	lines map[string][]*model.Line
}

func NewResolver(qe queryengine.QueryEngine) *Resolver {
	return &Resolver{
		queryEngine: qe,
		lines:       make(map[string][]*model.Line),
	}
}
