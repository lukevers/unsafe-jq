package jq

import (
	"errors"
	"reflect"
	"strings"
)

// Query allows us to try and safely access data through a potential unsafe
// object by giving a jq-like query string, and returning a slice of results.
// If there are any issues along the way, the results returned will be nil,
// and an error will be returned.
func Query(query string, unsafeData interface{}) (results []interface{}, err error) {
	parts := strings.Split(query, ".")
	if parts[0] == "[]" {
		if reflect.TypeOf(unsafeData).Kind() != reflect.Slice {
			return nil, errors.New("Can not index over non-slice")
		}

		for _, dataObject := range unsafeData.([]interface{}) {
			result, err := Query(strings.Join(parts[1:], "."), dataObject)
			if err != nil {
				return nil, err
			}

			results = append(results, result...)
		}

		return
	}

	if unsafeData == nil {
		return nil, errors.New("Nil data")
	}

	value := unsafeData.(map[string]interface{})[parts[0]]
	if len(parts) > 1 {
		return Query(strings.Join(parts[1:], "."), value)
	}

	results = append(results, value)
	return
}
