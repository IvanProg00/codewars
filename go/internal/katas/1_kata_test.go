package katas_test

import (
	"codewars/internal/katas"
	"testing"
)

func TestDNAStrand(t *testing.T) {
	type StructTest struct {
		value    string
		expected string
	}

	tests := []StructTest{
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
			t.Errorf("%s is not equals to %s", res, test.expected)
		}
	}
}
