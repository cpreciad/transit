package main

import (
	"fmt"
	"log/slog"
	"os"

	apiqe "github.com/cpreciad/transit/internal/api_query_engine"
)

const apiKeyEnv = "TRANSIT_DATA_API_KEY"

// this is just to test the api query engine implementation
func main() {
	apiKey := os.Getenv(apiKeyEnv)
	if apiKey == "" {
		panicMessage := fmt.Sprintf("NewApiQueryEngine: an api key is not mapped to the env variable %s. Please go to 511.org, register for an api key, and set it to the listed env variable", apiKeyEnv)
		panic(panicMessage)
	}
	api := apiqe.NewApiQueryEngine(apiKey)
	outO, err := api.GetOperatorID()

	if err != nil {
		slog.Error("ez error", "returned error:", err.Error())
	}
	fmt.Println(outO)

	for _, v := range outO {
		l, err := api.GetLineID(v.OperatorID)
		if err != nil {
			slog.Error("ez error", "returned error:", err.Error())
		}
		fmt.Println(v.OperatorID)
		fmt.Println(l)
		fmt.Println("----")

	}

}
