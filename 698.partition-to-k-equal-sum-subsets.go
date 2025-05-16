package neetcode250

import "slices"

/*
 * @lc app=leetcode id=698 lang=golang
 *
 * [698] Partition to K Equal Sum Subsets
 */

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {

	// var partitionSum []int = make([]int, k)

	slices.Sort(nums)

	var dfs func(i int)
	dfs = func(i int) {

	}

	dfs(0)

	return false
}

// @lc code=end
