package neetcode250

/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	var left = 0
	var window int          // for sum
	result := len(nums) + 1 // initialize to a value larger than possible

	// invalid : window is invalid if sum < target
	isValid := func() bool {
		return window < target
	}

	// Define process function
	process := func(right int) {
		result = min(result, right-left+1)
	}

	for right := 0; right < len(nums); right++ {
		// expand the window to the right
		window += nums[right]

		// while the current window is valid (sum >= target)
		for !isValid() {
			// now the window is valid; update minimum length
			process(right)

			// shrink the window from the left
			window -= nums[left]
			left++
		}
	}

	if result > len(nums) {
		return 0
	}

	return result
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
