package katas

import "fmt"

// Execrsice from interview: Find max sum
func findMaxSum(numbers []int) int {
	first := true
	bigger := 0
	for i := 0; i < len(numbers)-1; i++ {
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
	fmt.Println(findMaxSum(list))
}
