package katas

import (
	"fmt"
	"math"
)

func RunKata2() {
	fmt.Println(isPrimeNumber(5))
}

func isPrimeNumber(num int) bool {
	if num <= 3 {
		return num > 1
	}

	if num%2 == 0 || num%3 == 0 {
		return false
	}

	i := 5
	for int(math.Pow(float64(i), 2)) <= num {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}
