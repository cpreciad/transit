package request

import (
    "net/http"
    "fmt"
    "io"
    "os"
    "log"
)

const(
    transitUrl = "http://api.511.org/transit/stops"
    apiKeyEnv = "TRANSIT_DATA_API_KEY"
    operatorID = "SF"
    lineID = "N"
)


// RequestStops - returns a byte slice of all stops along the N line
func RequestStops() ([]byte, error) {
    url, err := requestStopsConstructUrl()
    if err != nil{
        return nil, err
    }

    resp, err := http.Get(url)
    if err != nil{
        return nil, fmt.Errorf("Request: HTTP protocol error  %s\n", url)
    }
    if resp.StatusCode != 200{
        return nil, fmt.Errorf("Request: bad status code %d\n", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)

    // clean up the response body
    body = clean(body)

    resp.Body.Close()

    return body, nil

}

// RequestNextArrivals - takes a stop ID and returns a byte slice of next arrivals 
//                       for the specified stop
func RequestNextArrivals() ([]byte, error){
    return nil, nil
}

func requestStopsConstructUrl() (string, error){
    apiKey, err := requestGetApiKey()
    if err != nil{
        return "", err
    }

    url := fmt.Sprintf("%s?api_key=%s&operator_id=%s&line_id=%s&format=json",transitUrl, apiKey, operatorID, lineID)
    log.Printf("Request: constructed url %s\n", url)

   return url, nil

}

func requestGetApiKey() (string, error){
    value := os.Getenv(apiKeyEnv)
    if value == ""{
        return "", fmt.Errorf("Request: %s has not been set", apiKeyEnv)
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
        b[2] == 0xbf{
            return b[3:]
        }
    return b
}
