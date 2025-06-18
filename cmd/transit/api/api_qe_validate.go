package main

import (
	"fmt"
	"log/slog"

	"github.com/cpreciad/transit/internal"
)

// this is just to test the api query engine implementation
func main() {
	apiqe := internal.NewApiQueryEngine()
	out, err := apiqe.GetOperatorID()
	if err != nil {
		slog.Error("ez error", "returned error:", err.Error())
	}
	fmt.Println(out)
}
