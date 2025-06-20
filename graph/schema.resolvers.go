package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"log/slog"
	"sort"
	"strconv"

	"github.com/cpreciad/transit/cmd/transit/duboce/consolidator"
	"github.com/cpreciad/transit/graph/model"
	qe "github.com/cpreciad/transit/query_engine"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Operators is the resolver for the operators field.
func (r *queryResolver) Operators(ctx context.Context, order *model.SortOrder) ([]*model.Operator, error) {
	operators, err := r.QueryEngine.GetOperatorID()
	if err != nil {
		slog.Error("QueryResolver", "Method", "Operator", "Error", err.Error())
		return nil, gqlerror.Errorf("Internal server error occurred")
	}
	sortedKeys := make([]qe.ID, 0, len(operators))
	for i := range operators {
		sortedKeys = append(sortedKeys, i)
	}

	sort.Slice(sortedKeys, func(i, j int) bool {
		var less bool

		l, _ := strconv.Atoi(string(sortedKeys[i]))
		r, _ := strconv.Atoi(string(sortedKeys[j]))

		less = l > r

		if *order == model.SortOrderAsc {
			less = !less

		}
		return less
	})

	for _, opID := range sortedKeys {
		op := operators[opID]
		slog.Info("op info: ", "id", string(opID), "opID", op.OperatorID, "name", op.Name)
		r.operators = append(r.operators, &model.Operator{
			ID:         string(opID),
			OperatorID: op.OperatorID,
			Name:       op.Name,
			Lines:      nil, // TODO: implement recursive calls to Lines
		})
	}

	return r.operators, nil
}

// Operator is the resolver for the operator field.
func (r *queryResolver) Operator(ctx context.Context, id string) (*model.Operator, error) {
	operators, err := r.QueryEngine.GetOperatorID()
	if err != nil {
		slog.Error("QueryResolver", "Method", "Operator", "Error", err.Error())
		return nil, gqlerror.Errorf("Internal server error occurred")
	}

	operator, ok := operators[qe.ID(id)]
	if !ok {
		return nil, gqlerror.Errorf("Operator with ID %s could not be found", id)
	}

	return &model.Operator{
		ID:         operator.ID,
		OperatorID: operator.OperatorID,
		Name:       operator.Name,
		Lines:      nil, // TODO: implement recursive calls to Lines
	}, nil
}

// StopsForLine is the resolver for the stopsForLine field.
func (r *queryResolver) StopsForLine(ctx context.Context, operatorID string, lineID string) ([]*model.Stop, error) {
	stops := make(map[string][]string)
	stops["Carl St & Cole St"] = make([]string, 0)
	stops["Duboce St/Noe St/Duboce Park"] = make([]string, 0)
	stopInfo := consolidator.GetStopInfo(operatorID, lineID, stops)

	// construct the return for the stops
	for _, stop := range stopInfo {
		formattedStop := &model.Stop{
			ID:   "1",
			Name: stop.Direction.Outbound.StopName,
		}
		r.stops = append(r.stops, formattedStop)
	}
	return r.stops, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
