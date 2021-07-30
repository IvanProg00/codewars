package katas

import (
	"fmt"
	"strings"
)

const PosibleLetters = "abcdefghijklm"

// Printer Errors: https://www.codewars.com/kata/56541980fa08ab47a0000040/train/go
func RunKata5() {
	fmt.Println(PrinterError("aaabbbbhaijjjm"))
	fmt.Println(PrinterError("aaaxbbbbyyhwawiwjjjwwm"))
}

func PrinterError(s string) string {
	size := len(s)
	errors := 0

	for _, c := range s {
		if !strings.ContainsRune(PosibleLetters, c) {
			errors++
		}
	}

	return fmt.Sprintf("%d/%d", errors, size)
}
