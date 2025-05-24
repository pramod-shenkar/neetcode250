package neetcode250

/*
 * @lc app=leetcode id=424 lang=golang
 *
 * [424] Longest Repeating Character Replacement
 */

// @lc code=start
func characterReplacement(s string, k int) int {

	var left, right = 0, 0
	var window = map[byte]int{}
	var max1 int

	// we need to replace how many chars
	// max chodke jitne bhi char hai saare replace karna hai
	// ie need_to_replace = window_size - max_char_count
	var isValid = func(i int, windowSize int) bool {

		// find max for that char & update if needed
		max1 = max(max1, window[s[i]])
		return windowSize-max1 <= k
	}

	var result int
	var process = func(currentWindowLength int) {
		result = max(result, currentWindowLength)
	}

	for right = 0; right < len(s); right++ {

		window[s[right]]++

		// for invalid case
		if !isValid(right, right-left+1) {
			// move window
			window[s[left]]--
			left++
		}

		// for vaid case
		process(right - left + 1)

	}

	return result

}

// @lc code=end
