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
	linesUrl = "http://api.511.org/transit/lines"
)

type Line struct {
	LineID string `json:"Id"`
	Name   string `json:"Name"`
}

func lineFetch(apiKey, oid string) ([]byte, error) {
	url := fmt.Sprintf("%s?api_key=%s&operator_id=%s&format=json", linesUrl, apiKey, oid)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("lineFetch: bad status code returned: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("lineFetch: %s", err.Error())
	}
	body = helpers.CleanResponseBody(body)
	resp.Body.Close()

	return body, nil

}
func lineFormat(data []byte) (map[qe.ID]qe.Line, error) {
	var unpackedLines []Line
	if err := json.Unmarshal(data, &unpackedLines); err != nil {
		return nil, fmt.Errorf("lineFormat: %w", err)
	}

	m := make(map[qe.ID]qe.Line)
	for i, l := range unpackedLines {
		// getting in the habit of using database ids, as this data will eventually be stored there
		id := qe.ID(strconv.Itoa(i + 1))
		m[id] = qe.Line{
			ID:     string(id),
			LineID: l.LineID,
			Name:   l.Name,
		}
	}

	return m, nil
}
func lineValidate(map[qe.ID]qe.Line) error {
	return nil
}

// A SAMPLE OF WHAT THE DATA LOOKS LIKE
// values can either be strings or null, likely transposing to the empty string ""
/*
{"Id":"49",
"Name":"VAN NESS-MISSION",
"FromDate":"2025-06-21T00:00:00-07:00",
"ToDate":"2025-08-29T23:59:00-07:00",
"TransportMode":"bus",
"PublicCode":"49",
"SiriLineRef":"49",
"Monitored":true,
"OperatorRef":"SF"}
*/
