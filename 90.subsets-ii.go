package neetcode250

import (
	"slices"
)

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {

	var result [][]int

	slices.Sort(nums)

	var dfs func(i int, subset []int)
	dfs = func(i int, subset []int) {

		if i > len(nums) {
			return
		}

		var tempset = make([]int, len(subset))
		copy(tempset, subset)
		result = append(result, tempset)

		for j := i; j < len(nums); j++ {

			if j > i && nums[j] == nums[j-1] {
				continue
			}

			subset = append(subset, nums[j])
			dfs(j+1, subset)
			subset = subset[:len(subset)-1]
		}
	}

	dfs(0, nil)

	return result
}

// @lc code=end

/*
	1		2
	2	2	2
	2


*/
