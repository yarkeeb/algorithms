package string

// leetcode 125
// https://leetcode.com/problems/valid-palindrome/

import "strings"
import "unicode"

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if !isAlphanumeric(rune(s[i])) {
			i++
			continue
		}
		if !isAlphanumeric(rune(s[j])) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
