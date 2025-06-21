package main

import (
	"fmt"
	"log/slog"

	apiqe "github.com/cpreciad/transit/internal/api_query_engine"
)

// this is just to test the api query engine implementation
func main() {
	apiqe := apiqe.NewApiQueryEngine()
	out, err := apiqe.GetOperatorID()
	if err != nil {
		slog.Error("ez error", "returned error:", err.Error())
	}
	fmt.Println(out)
}
