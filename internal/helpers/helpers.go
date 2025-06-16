package helpers

import (
	"time"
)

func UTCtoPST(utcTime string) (time.Time, error) {

	pst, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return time.Time{}, err
	}

	t, err := time.Parse(time.RFC3339, utcTime)
	if err != nil {
		return time.Time{}, err
	}
	t = t.In(pst)

	return t, nil
}

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
