package main

import (
	"reflect"
	"testing"
)

func TestExtractPrNumber(t *testing.T) {
	type test struct {
		input string
		exp   string
	}

	tests := []test{
		{input: "github.com/shayansadeghieh/foo/bar/1", exp: "1"},
		{input: "github.com/shayansadeghieh/random/10", exp: "10"},
		{input: "github.com/shayansadeghieh/random/abc", exp: "abc"},
	}

	for _, tc := range tests {
		got := extractPrNumber(tc.input)
		if !reflect.DeepEqual(tc.exp, got) {
			t.Fatalf("expected: %v, got: %v", tc.exp, got)
		}
	}
}
