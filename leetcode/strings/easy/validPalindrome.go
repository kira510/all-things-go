package main

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if !isAlfaNumeric(s[i]) {
			i++
			continue
		} else if !isAlfaNumeric(s[j]) {
			j--
			continue
		}

		if isLower(s[i]) != isLower(s[j]) {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlfaNumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

func isLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}

	return b
}

// https://leetcode.com/problems/valid-palindrome/description/
