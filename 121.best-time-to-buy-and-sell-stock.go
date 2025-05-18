package neetcode250

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {

	var result int
	var left = 0

	// expand window in for iteself
	for right := 0; right < len(prices); right++ {

		currentProfit := prices[right] - prices[left]

		// if window is invalid
		if currentProfit <= 0 {
			// move window
			left = right

		}

		// if window is valid
		if currentProfit > 0 {
			// process
			result = max(result, currentProfit)
		}

	}

	return result
}

// @lc code=end
