package neetcode250

/*
 * @lc app=leetcode id=1863 lang=golang
 *
 * [1863] Sum of All Subset XOR Totals
 */

// @lc code=start
func subsetXORSum(nums []int) int {

	var result = 0

	var xor = 0

	var dfs func(int)
	dfs = func(i int) {

		if i == len(nums) {
			result += xor
			return
		}

		// choose element
		dfs(i + 1)

		// not choose element
		xor = xor ^ nums[i] // why we done same??? We need to undo. And in xor 2^4=6 then undo is 6^4 = 2
		dfs(i + 1)

	}

	dfs(0)

	return result
}

// @lc code=end
