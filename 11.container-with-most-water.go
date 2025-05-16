package neetcode250

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {

	var result int

	var start, end = 0, len(height) - 1

	for start < end {

		m := min(height[start], height[end])
		distance := end - start
		ans := m * distance

		result = max(ans, result)

		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}

	return result
}

// @lc code=end
