package internal

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/cpreciad/transit/internal/helpers"
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
		return nil, fmt.Errorf("apiQueryEngine: fetchData: bad status code returned: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("apiQueryEngine: fetchData: %s", err.Error())
	}
	body = helpers.CleanResponseBody(body)
	resp.Body.Close()

	return body, nil

}

func formatApiData(body []byte) (map[qe.ID]qe.Operator, error) {
	fmt.Println(string(body))
	// unpack the go object from the operator

	// validate the unpacked object

	// iterate through the go object and format the API data to fit the expected interface return
	return nil, nil
}

// checks if an API key exists as an error check
func constructOperatorsUrl() (string, error) {
	apiKey := os.Getenv(apiKeyEnv)
	if apiKey == "" {
		return "", fmt.Errorf("apiQueryEngine: constructOperatorsUrl: %s env variable is not set", apiKeyEnv)
	}

	url := fmt.Sprintf("%s?api_key=%s&format=json", operatorsUrl, apiKey)

	return url, nil
}
