package katas_test

import (
	"codewars/internal/katas"
	"testing"
)

func TestLongestConsec(t *testing.T) {
	type ValueStruct struct {
		strarr []string
		k      int
	}
	type StructTest struct {
		value    ValueStruct
		expected string
	}
	tests := []StructTest{
		{
			value: ValueStruct{
				strarr: []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},
				k:      1,
			},
			expected: "oocccffuucccjjjkkkjyyyeehh",
		},
	}

	for _, test := range tests {
		res := katas.LongestConsec(test.value.strarr, test.value.k)
		if res != test.expected {
			t.Errorf("%s is not equals to %s", res, test.expected)
		}
	}
}
