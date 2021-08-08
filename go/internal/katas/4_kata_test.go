package katas_test

import (
	"codewars/internal/katas"
	"testing"
)

func TestScale(t *testing.T) {
	type Value struct {
		S string
		K int
		N int
	}
	type TestStruct struct {
		Value    Value
		Expected string
	}

	tests := []TestStruct{
		{
			Value: Value{
				S: "abcd\nefgh\nijkl\nmnop",
				K: 2,
				N: 3,
			},
			Expected: "aabbccdd\naabbccdd\naabbccdd\neeffgghh\neeffgghh\neeffgghh\niijjkkll\niijjkkll\niijjkkll\nmmnnoopp\nmmnnoopp\nmmnnoopp",
		},
	}

	for _, test := range tests {
		res := katas.Scale(test.Value.S, test.Value.K, test.Value.N)
		if res != test.Expected {
			t.Errorf("%s is not equals to %s", res, test.Expected)
		}
	}
}
