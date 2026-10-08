package medium

import "strings"

func reverseWords(s string) string {
	var result []string

	i := len(s) - 1
	for i >= 0 {
		for i >= 0 && s[i] == ' ' {
			i--
		}
		if i < 0 {
			break
		}

		j := i
		for j >= 0 && s[j] != ' ' {
			j--
		}

		result = append(result, s[j+1:i+1])
		i = j
	}

	return strings.Join(result, " ")
}

// https://leetcode.com/problems/reverse-words-in-a-string/
