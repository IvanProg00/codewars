package katas_test

import (
	"codewars/internal/katas"
	"testing"
)

func TestDNAStrand(t *testing.T) {
	type StructTest struct {
		Value    string
		Expected string
	}

	tests := []StructTest{
		{
			Value:    "ATTGC",
			Expected: "TAACG",
		}, {
			Value:    "GTAT",
			Expected: "CATA",
		},
	}

	for _, test := range tests {
		res := katas.DNAStrand(test.Value)
		if res != test.Expected {
			t.Errorf("%s is not equals to %s", res, test.Expected)
		}
	}
}
