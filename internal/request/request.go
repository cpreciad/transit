package request

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	apiKeyEnv = "TRANSIT_DATA_API_KEY"

	stopsUrl          = "http://api.511.org/transit/stops"
	stopMonitoringUrl = "http://api.511.org/transit/StopMonitoring"
)

// RequestStops - returns a byte slice of all stops along the N line
func RequestStops(operatorId, lineId string) ([]byte, error) {
	url, err := requestStopsConstructUrl(operatorId, lineId)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request: bad status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// clean up the response body
	body = clean(body)

	resp.Body.Close()

	return body, nil

}

// RequestNextArrivals - takes a stop ID and returns a byte slice of next arrivals
//
//	for the specified stop
func RequestNextArrivals(operatorId, stopId string) ([]byte, error) {
	url, err := requestStopMonitoringConstructUrl(operatorId, stopId)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request: bad status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body = clean(body)
	resp.Body.Close()

	return body, nil
}

func requestStopsConstructUrl(operatorId, lineId string) (string, error) {
	apiKey, err := requestGetApiKey()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s?api_key=%s&operator_id=%s&line_id=%s&format=json", stopsUrl, apiKey, operatorId, lineId)

	return url, nil

}

func requestStopMonitoringConstructUrl(operatorId, stopCode string) (string, error) {
	apiKey, err := requestGetApiKey()
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s?api_key=%s&agency=%s&stopCode=%s&format=json", stopMonitoringUrl, apiKey, operatorId, stopCode)

	return url, nil

}

func requestGetApiKey() (string, error) {
	value := os.Getenv(apiKeyEnv)
	if value == "" {
		return "", fmt.Errorf("request: %s env variable has not been set", apiKeyEnv)
	}
	return value, nil
}

func clean(b []byte) []byte {
	// https://en.wikipedia.org/wiki/Byte_order_mark
	// check that the first three runes of the byte array are the Byte Order Mark
	// of UTF-8, and return a byte array that trims these off
	if len(b) >= 3 &&
		b[0] == 0xef &&
		b[1] == 0xbb &&
		b[2] == 0xbf {
		return b[3:]
	}
	return b
}
