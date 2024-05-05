package internal

import (
	"reflect"
	"testing"
)

func TestExtractNames(t *testing.T) {
	testCases := []struct {
		name    string
		members AllMembers
		exp     []string
	}{
		{
			name: "Multiple members",
			members: AllMembers{
				Members: []Member{
					{Profile: Profile{RealNameNormalized: "shaya sadeghieh"}},
					{Profile: Profile{RealNameNormalized: "foo bar"}},
					{Profile: Profile{RealNameNormalized: "foo"}},
				},
			},
			exp: []string{"shayan sadeghieh", "foo bar", "foo"},
		},
		{
			name: "No members",
			members: AllMembers{
				Members: []Member{},
			},
			exp: []string{},
		},
	}

	// Iterate over test cases and run them
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Run the function being tested
			names := extractNames(tc.members)
			// Check if the output matches the expected result
			if !reflect.DeepEqual(names, tc.exp) {
				t.Errorf("Expected %v, but got %v", tc.exp, names)
			}
		})
	}

}
