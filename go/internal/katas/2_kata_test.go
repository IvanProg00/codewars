package katas_test

import (
	"codewars/internal/katas"
	"testing"
)

type StructTest struct {
	value    int
	expected bool
}

func TestIsPrimeNumber(t *testing.T) {
	tests := []StructTest{
		{
			value:    5,
			expected: true,
		},
		{
			value:    150,
			expected: false,
		},
		{
			value:    385,
			expected: false,
		},
		{
			value:    401,
			expected: true,
		},
		{
			value:    577,
			expected: true,
		},
	}

	for _, test := range tests {
		if katas.IsPrimeNumber(test.value) != test.expected {
			if test.expected {
				t.Errorf("%d must be a primitive number", test.value)
			} else {
				t.Errorf("%d must not be a primitive number", test.value)
			}
		}
	}
}
