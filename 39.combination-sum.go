package neetcode250

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start

func combinationSum(nums []int, target int) [][]int {

	var result [][]int

	var dfs func(int, int, []int)
	dfs = func(i int, sum int, subset []int) {
		if sum == target {
			var tempset = make([]int, len(subset))
			copy(tempset, subset)
			result = append(result, tempset)
			return
		}

		if sum > target {
			return
		}

		for j := i; j < len(nums); j++ {

			subset = append(subset, nums[j])
			dfs(j, sum+nums[j], subset)
			subset = subset[:len(subset)-1]
		}

	}

	dfs(0, 0, []int{})

	return result
}

// @lc code=end
