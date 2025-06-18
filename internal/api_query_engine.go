package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/cpreciad/transit/internal/helpers"
	"github.com/cpreciad/transit/internal/model"
	qe "github.com/cpreciad/transit/query_engine"
)

const (
	apiKeyEnv    = "TRANSIT_DATA_API_KEY"
	operatorsUrl = "http://api.511.org/transit/operators"
)

type apiQueryEngine struct{}

func (a *apiQueryEngine) GetOperatorID() (map[qe.ID]qe.Operator, error) {

	// get the body of the request to the url
	body, err := fetchData()
	if err != nil {
		return nil, err
	}

	// given the body of the API response, format it into a map[id]Operator
	out, err := formatApiData(body)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func NewApiQueryEngine() *apiQueryEngine {
	return &apiQueryEngine{}
}

func fetchData() ([]byte, error) {
	url, err := constructOperatorsUrl()
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("fetchData: bad status code returned: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("fetchData: %s", err.Error())
	}
	body = helpers.CleanResponseBody(body)
	resp.Body.Close()

	return body, nil

}

func formatApiData(data []byte) (map[qe.ID]qe.Operator, error) {
	var unpackedOperators []model.Operator
	if err := json.Unmarshal(data, &unpackedOperators); err != nil {
		return nil, fmt.Errorf("formatApiData: %w", err)
	}

	if err := validate(unpackedOperators); err != nil {
		return nil, fmt.Errorf("formatApiData: %w", err)
	}

	m := make(map[qe.ID]qe.Operator)
	for i, o := range unpackedOperators {
		// getting in the habit of using database ids, as this data will eventually be stored there
		id := qe.ID(strconv.Itoa(i + 1))
		m[id] = qe.Operator{
			ID:         string(id),
			OperatorID: o.OperatorID,
			Name:       o.Name,
		}
	}

	return m, nil
}

// checks if an API key exists as an error check
func constructOperatorsUrl() (string, error) {
	apiKey := os.Getenv(apiKeyEnv)
	if apiKey == "" {
		return "", fmt.Errorf("constructOperatorsUrl: %s env variable is not set", apiKeyEnv)
	}

	url := fmt.Sprintf("%s?api_key=%s&format=json", operatorsUrl, apiKey)

	return url, nil
}

// TODO: find a way to validate the unpacking after a successful unmarshal
func validate([]model.Operator) error {
	return nil
}
