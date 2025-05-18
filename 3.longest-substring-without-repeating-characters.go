package neetcode250

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {

	var left = 0

	var window = map[byte]int{} // e.g., for count as we avoidng duplicate

	// valid current char shouldnt be exist in window to avoid duplicate
	var isValid = func(i int) bool {
		return window[s[i]] <= 1
	}

	var result int
	var process = func(cuurentWindowLength int) {
		result = max(result, cuurentWindowLength)
	}

	for right := 0; right < len(s); right++ {
		// expand the window to the ie add in count map
		window[s[right]]++

		// check current element is valid or not
		for !isValid(right) {
			// shrink the window from the left
			window[s[left]]--
			left++
		}

		// for valid window
		process(right - left + 1) // +1 just because we need convert index diff in legth
	}

	return result
}

// @lc code=end
