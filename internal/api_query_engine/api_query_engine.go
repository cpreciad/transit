package apiqe

import (
	qe "github.com/cpreciad/transit/query_engine"
)

const (
	apiKeyEnv = "TRANSIT_DATA_API_KEY"
)

type apiQueryEngine struct{}

func NewApiQueryEngine() *apiQueryEngine {
	return &apiQueryEngine{}
}

// TODO: figure out how to make factory functions for each transit type
func (a *apiQueryEngine) GetOperatorID() (map[qe.ID]qe.Operator, error) {

	out, err := fetchAndFormatData(operatorFetch, operatorFormat, operatorValidate)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (a *apiQueryEngine) GetLineID(oid string) (map[qe.ID]qe.Line, error) {
	// make the fetch function generic
	fetch := func() ([]byte, error) {
		return lineFetch(oid)
	}

	out, err := fetchAndFormatData(fetch, lineFormat, lineValidate)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func fetchAndFormatData[K comparable, V qe.Operator | qe.Line](fetch func() ([]byte, error), format func([]byte) (map[K]V, error), validate func(map[K]V) error) (map[K]V, error) {

	// get the body of the request to the 511 API
	body, err := fetch()
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

// TODO: find a way to validate the unpacking after a successful unmarshal
func validate([]Operator) error {
	return nil
}
