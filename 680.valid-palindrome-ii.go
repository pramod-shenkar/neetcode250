package neetcode250

/*
 * @lc app=leetcode id=680 lang=golang
 *
 * [680] Valid Palindrome II
 */

// @lc code=start
func validPalindrome(s string) bool {
	var start, end = 0, len(s) - 1

	var alreadyDeleted bool

	for start <= end {

		if s[start] != s[end] {

			if alreadyDeleted {
				return false
			} else {
				if s[start+1] == s[end] {
					start++
					alreadyDeleted = true
				} else if s[start] == s[end-1] {
					end--
					alreadyDeleted = true
				} else {
					return false
				}
			}

		} else {
			start++
			end--
		}

	}

	return true
}

// @lc code=end
