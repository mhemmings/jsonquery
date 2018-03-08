package jsonquery

import (
	"encoding/json"
	"errors"
	"net/url"
	"strings"
)

var ErrInvalidJSON = errors.New("The query string provided must be a JSON array of strings")

// HandleURL mutates the given URL.
// Any query string with given prefix is unmarshalled as a JSON array of strings.
// The results are added to url.RawQuery under a single query param (minus the prefix)
// Returns an error if a matched query is not a JSON string
// See tests for examples
func HandleURL(u *url.URL, prefix string) error {
	newQueries := u.Query()

	for key, _ := range u.Query() {
		if strings.HasPrefix(key, prefix) {
			newKey := strings.TrimPrefix(key, prefix)
			parsedVals, err := parseQuery(u.Query().Get(key))
			if err != nil {
				return err
			}

			for _, str := range parsedVals {
				newQueries.Add(newKey, str)
			}
		}
	}

	u.RawQuery = newQueries.Encode()

	return nil
}

func parseQuery(query string) ([]string, error) {
	var slice []string
	err := json.Unmarshal([]byte(query), &slice)
	if err != nil {
		return nil, ErrInvalidJSON
	}
	return slice, nil
}
