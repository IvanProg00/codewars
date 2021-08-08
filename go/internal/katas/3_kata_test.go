package katas_test

import (
	"codewars/internal/katas"
	"testing"
)

func TestLongestConsec(t *testing.T) {
	type ValueStruct struct {
		Strarr []string
		K      int
	}
	type StructTest struct {
		Value    ValueStruct
		Expected string
	}
	tests := []StructTest{
		{
			Value: ValueStruct{
				Strarr: []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},
				K:      1,
			},
			Expected: "oocccffuucccjjjkkkjyyyeehh",
		},
	}

	for _, test := range tests {
		res := katas.LongestConsec(test.Value.Strarr, test.Value.K)
		if res != test.Expected {
			t.Errorf("%s is not equals to %s", res, test.Expected)
		}
	}
}
