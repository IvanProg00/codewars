package katas_test

import (
	"codewars/internal/katas"
	"testing"
)

func TestPrinterError(t *testing.T) {
	type TestStruct struct {
		Value    string
		Expected string
	}

	tests := []TestStruct{
		{
			Value:    "aaabbbbhaijjjm",
			Expected: "0/14",
		},
		{
			Value:    "aaaxbbbbyyhwawiwjjjwwm",
			Expected: "8/22",
		},
	}

	for _, test := range tests {
		res := katas.PrinterError(test.Value)
		if res != test.Expected {
			t.Errorf("%s is not equals to %s", res, test.Expected)
		}
	}
}
