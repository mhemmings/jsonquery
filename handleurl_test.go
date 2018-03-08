package jsonquery

import (
	"net/url"
	"testing"
)

var testTable = []struct {
	URL    string
	Query  string
	Values []string
	Err    error
}{
	{
		URL:    `https://example.com?json_foo=["bar", "baz"]`,
		Query:  "foo",
		Values: []string{"bar", "baz"},
	},
	{
		URL:   `https://example.com?json_foo=[bar", "baz"]`,
		Query: "foo",
		Err:   ErrInvalidJSON,
	},
}

func TestParseURL(t *testing.T) {
	for i, test := range testTable {
		u, err := url.Parse(test.URL)
		if err != nil {
			t.Fatalf("[%d] Error parsing URL: %s", i, err)
		}

		err = HandleURL(u, "json_")
		if err != test.Err {
			t.Fatalf("[%d]  Error handling URL: %s", i, err)
		}

		foos := u.Query()[test.Query]

		equal := equalSlices(foos, test.Values)
		if !equal {
			t.Fatalf("[%d] Slices do not match: Exected: %v, got: %v", i, test.Values, foos)
		}
	}
}

// https://stackoverflow.com/questions/36000487/check-for-equality-on-slices-without-order
func equalSlices(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}
	for _, _y := range y {
		// If the string _y is not in diff bail out early
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}
