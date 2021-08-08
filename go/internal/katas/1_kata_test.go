package katas_test

import (
	"codewars/internal/katas"
	"testing"
)

type expectedTest struct {
	value    string
	expected string
}

func TestDNAStrand(t *testing.T) {
	tests := []expectedTest{
		{
			value:    "ATTGC",
			expected: "TAACG",
		}, {
			value:    "GTAT",
			expected: "CATA",
		},
	}

	for _, test := range tests {
		res := katas.DNAStrand(test.value)
		if res != test.expected {
			t.Errorf("%s not equals to %s", res, test.expected)
		}
	}
}
