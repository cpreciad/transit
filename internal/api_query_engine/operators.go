package apiqe

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/cpreciad/transit/internal/helpers"
	qe "github.com/cpreciad/transit/query_engine"
)

const (
	operatorsUrl = "http://api.511.org/transit/operators"
)

// these struct fields need to be exported so the json lib sees them, idk
type Operator struct {
	OperatorID string `json:"Id"`
	Name       string `json:"Name"`
}

func operatorFetch(apiKey string) ([]byte, error) {

	url := fmt.Sprintf("%s?api_key=%s&format=json", operatorsUrl, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("operatorFetch: bad status code returned: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("operatorFetch: %s", err.Error())
	}
	body = helpers.CleanResponseBody(body)
	resp.Body.Close()

	return body, nil
}

func operatorFormat(data []byte) (map[qe.ID]qe.Operator, error) {
	var unpackedOperators []Operator
	if err := json.Unmarshal(data, &unpackedOperators); err != nil {
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

// make sure all of the data we care about is here, and return an error explaining what is missing
func operatorValidate(map[qe.ID]qe.Operator) error {
	return nil
}

// A SAMPLE OF WHAT THE DATA LOOKS LIKE
// values can either be strings or null, likely transposing to the empty string ""
/*
{
    "Id": "SS",
    "Name": "City of South San Francisco",
    "ShortName": "South City Shuttle",
    "SiriOperatorRef": null,
    "TimeZone": "America/Vancouver",
    "DefaultLanguage": "en",
    "ContactTelephoneNumber": null,
    "WebSite": "http://www.ssf.net/SCS",
    "PrimaryMode": "bus",
    "PrivateCode": "SS",
    "Monitored": false,
    "OtherModes": ""
  },
*/
