package katas

import (
	"fmt"
	"strings"
)

func RunKata4() {
	res1 := Scale("abcd\nefgh\nijkl\nmnop", 2, 3)
	expected1 := "aabbccdd\naabbccdd\naabbccdd\neeffgghh\neeffgghh\neeffgghh\niijjkkll\niijjkkll\niijjkkll\nmmnnoopp\nmmnnoopp\nmmnnoopp"
	fmt.Println(res1)
	fmt.Println(expected1)
	fmt.Println(res1 == expected1)
	fmt.Println(Scale("", 4, 4))
}

func Scale(s string, k, n int) string {
	if s == "" {
		return ""
	}

	strs := strings.Split(s, "\n")
	res := []string{}

	for _, str := range strs {
		appendStr := ""
		for _, c := range str {
			appendStr += strings.Repeat(string(c), k)
		}
		for i := 0; i < n; i++ {
			res = append(res, appendStr)
		}
	}

	return strings.Join(res, "\n")
}
