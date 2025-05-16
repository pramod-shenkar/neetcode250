package neetcode250

import "unicode"

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {

	var start, end int = 0, len(s) - 1

	for start <= end {

		if !IsChar(s[start]) {
			start++
			continue
		}

		if !IsChar(s[end]) {
			end--
			continue
		}

		if ToLower(s[start]) != ToLower(s[end]) {
			return false
		}

		start++
		end--
	}

	return true
}

func IsChar(s byte) bool {
	return unicode.IsLetter(rune(s)) || unicode.IsDigit(rune(s))
}

func ToLower(s byte) byte {
	return byte(unicode.ToLower(rune(s)))
}

// @lc code=end
