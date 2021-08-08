package katas_test

import (
	"codewars/internal/katas"
	"testing"
)

func TestFindMaxSum(t *testing.T) {
	type TestStruct struct {
		Value    []int
		Expected int
	}
	tests := []TestStruct{
		{
			Value:    []int{5, 9, 7, 11},
			Expected: 20,
		},
	}

	for _, test := range tests {
		res := katas.FindMaxSum(test.Value)
		if res != test.Expected {
			t.Errorf("%d is not equals to %d", res, test.Expected)
		}
	}
}
