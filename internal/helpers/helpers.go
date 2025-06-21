package helpers

import (
	"fmt"
	"os"
)

func CleanResponseBody(b []byte) []byte {
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

// checks if an API key exists as an error check
func ConstructUrl(apiKeyEnv, url string) (string, error) {
	apiKey := os.Getenv(apiKeyEnv)
	if apiKey == "" {
		return "", fmt.Errorf("ConstructUrl: %s env variable is not set", apiKeyEnv)
	}

	constructedUrl := fmt.Sprintf("%s?api_key=%s&format=json", url, apiKey)

	return constructedUrl, nil
}
