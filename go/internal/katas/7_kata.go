package katas

import "fmt"

// Execrsice from interview: Find max sum
func FindMaxSum(numbers []int) int {
	first := true
	bigger := 0
	for i := range numbers {
		for k := i + 1; k < len(numbers); k++ {
			if i == k {
				continue
			}
			sum := numbers[i] + numbers[k]
			if first {
				first = false
				bigger = sum
			} else if sum > bigger {
				bigger = sum
			}
		}
	}

	return bigger
}

func RunKata7() {
	list := []int{5, 9, 7, 11}
	fmt.Println(FindMaxSum(list))
}
