package internal

import (
	"reflect"
	"testing"
)

func TestFilterNames(t *testing.T) {
	testCases := []struct {
		name   string
		names  []string
		filter string
		exp    []string
	}{
		{
			name:   "Multiple names returned",
			names:  []string{"shayan sadeghieh", "foo bar", "foo"},
			filter: "o",
			exp:    []string{"foo bar", "foo"},
		},
		{
			name:   "One name returned",
			names:  []string{"shayan sadeghieh", "foo bar", "foo"},
			filter: "shayan",
			exp:    []string{"shayan sadeghieh"},
		},
	}

	// Iterate over test cases and run them
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Run the function being tested
			filteredNames := filterNames(tc.names, tc.filter)

			if !reflect.DeepEqual(filteredNames, tc.exp) {
				t.Errorf("Mapping names to ID: Expected %v, but got %v", tc.exp, filteredNames)
			}

		})
	}
}
