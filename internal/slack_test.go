package internal

import (
	"reflect"
	"testing"
)

func TestMapNameToID(t *testing.T) {
	testCases := []struct {
		name     string
		members  AllMembers
		expMap   map[string]string
		expSlice []string
	}{
		{
			name: "Multiple members",
			members: AllMembers{
				Members: []Member{
					{Profile: Profile{RealNameNormalized: "shayan sadeghieh"},
						ID: "12345"},
					{Profile: Profile{RealNameNormalized: "foo bar"},
						ID: "abcd"},
					{Profile: Profile{RealNameNormalized: "foo"},
						ID: "abcd12345"},
				},
			},
			expMap: map[string]string{
				"shayan sadeghieh": "12345",
				"foo bar":          "abcd",
				"foo":              "abcd12345",
			},
			expSlice: []string{"shayan sadeghieh", "foo bar", "foo"},
		},
	}

	// Iterate over test cases and run them
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Run the function being tested
			namesToID, names := mapNameToID(tc.members)

			if !reflect.DeepEqual(namesToID, tc.expMap) {
				t.Errorf("Mapping names to ID: Expected %v, but got %v", tc.expMap, namesToID)
			}

			if !reflect.DeepEqual(names, tc.expSlice) {
				t.Errorf("Extracting names: Expected %v, but got %v", tc.expSlice, names)
			}
		})
	}
}
