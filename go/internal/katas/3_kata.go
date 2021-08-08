package katas

import (
	"strings"
)

func LongestConsec(strarr []string, k int) string {
	if k > len(strarr) || k <= 0 {
		return ""
	}

	max_len_str := ""

	for i := 0; i <= len(strarr)-k; i++ {
		str := strings.Join(strarr[i:i+k], "")

		if len(str) > len(max_len_str) {
			max_len_str = str
		}
	}

	return max_len_str
}
