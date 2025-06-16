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
	stops       []*model.Stop
	operators   []*model.Operator
	// the query engine interface should be injected here
}
