package katas_test

import (
	"codewars/internal/katas"
	"testing"
)

func TestIsPrimeNumber(t *testing.T) {
	type StructTest struct {
		value    int
		expected bool
	}

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
				t.Errorf("%d is not a primitive number", test.value)
			} else {
				t.Errorf("%d is a primitive number", test.value)
			}
		}
	}
}
