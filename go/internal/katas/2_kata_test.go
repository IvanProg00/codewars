package katas_test

import (
	"codewars/internal/katas"
	"testing"
)

func TestIsPrimeNumber(t *testing.T) {
	type StructTest struct {
		Value    int
		Expected bool
	}

	tests := []StructTest{
		{
			Value:    5,
			Expected: true,
		},
		{
			Value:    150,
			Expected: false,
		},
		{
			Value:    385,
			Expected: false,
		},
		{
			Value:    401,
			Expected: true,
		},
		{
			Value:    577,
			Expected: true,
		},
	}

	for _, test := range tests {
		if katas.IsPrimeNumber(test.Value) != test.Expected {
			if test.Expected {
				t.Errorf("%d is not a primitive number", test.Value)
			} else {
				t.Errorf("%d is a primitive number", test.Value)
			}
		}
	}
}
