package katas_test

import (
	"codewars/internal/katas"
	"testing"
)

func TestMovie(t *testing.T) {
	type Value struct {
		Card    int
		Ticket  int
		Percent float64
	}
	type TestStruct struct {
		Value    Value
		Expected int
	}

	tests := []TestStruct{
		{
			Value: Value{
				// fmt.Println(Movie(500, 15, 0.9))  // 43
				// fmt.Println(Movie(0, 10, 0.95))   // 2
				// fmt.Println(Movie(2500, 20, 0.9)) // 135
				Card:    500,
				Ticket:  15,
				Percent: 0.9,
			},
			Expected: 43,
		},
		{
			Value: Value{
				Card:    0,
				Ticket:  10,
				Percent: 0.95,
			},
			Expected: 2,
		},
		{
			Value: Value{
				Card:    2500,
				Ticket:  20,
				Percent: 0.9,
			},
			Expected: 135,
		},
	}

	for _, test := range tests {
		res := katas.Movie(test.Value.Card, test.Value.Ticket, test.Value.Percent)
		if res != test.Expected {
			t.Errorf("%d is not equals to %d", res, test.Expected)
		}
	}
}
