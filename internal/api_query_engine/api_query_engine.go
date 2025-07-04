package apiqe

import (
	"fmt"
	"os"

	qe "github.com/cpreciad/transit/query_engine"
)

const (
	apiKeyEnv = "TRANSIT_DATA_API_KEY"
)

type apiQueryEngine struct {
	apiKey string
}

func NewApiQueryEngine() *apiQueryEngine {
	apiKey := os.Getenv(apiKeyEnv)
	if apiKey == "" {
		panicMessage := fmt.Sprintf("NewApiQueryEngine: an api key is not mapped to the env variable %s. Please go to 511.org, register for an api key, and set it to the listed env variable", apiKeyEnv)
		panic(panicMessage)
	}
	return &apiQueryEngine{
		apiKey: apiKey,
	}
}

func (a *apiQueryEngine) GetOperatorID() (map[qe.ID]qe.Operator, error) {

	out, err := fetchAndFormatData(operatorFetch, operatorFormat, operatorValidate, a.apiKey)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (a *apiQueryEngine) GetLineID(oid string) (map[qe.ID]qe.Line, error) {
	// make the fetch function generic
	fetch := func(string) ([]byte, error) {
		return lineFetch(a.apiKey, oid)
	}

	out, err := fetchAndFormatData(fetch, lineFormat, lineValidate, a.apiKey)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func fetchAndFormatData[K comparable, V qe.Operator | qe.Line](fetch func(string) ([]byte, error), format func([]byte) (map[K]V, error), validate func(map[K]V) error, apiKey string) (map[K]V, error) {

	// get the body of the request to the 511 API
	body, err := fetch(apiKey)
	if err != nil {
		return nil, err
	}

	// given the body of the API response, format it into a map[id]Operator
	out, err := format(body)
	if err != nil {
		return nil, err
	}

	// check that the output is valid
	if err := validate(out); err != nil {
		return nil, err
	}

	return out, nil
}
