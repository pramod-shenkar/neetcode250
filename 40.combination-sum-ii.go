package neetcode250

import (
	"sort"
)

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {

	var result [][]int

	sort.Ints(candidates) // Sort first to handle duplicates

	var dfs func(i int, sum int, subset []int, depth int)
	dfs = func(i, sum int, subset []int, depth int) {

		if sum > target {
			return
		}

		// fmt.Printf("%sTrying %v, sum: %d\n", strings.Repeat("    ", depth), subset, sum)

		if sum == target {
			var tempset = make([]int, len(subset))
			copy(tempset, subset)
			result = append(result, tempset)
			return
		}

		for j := i; j < len(candidates); j++ {

			if j > i && candidates[j] == candidates[j-1] {
				continue
			}

			subset = append(subset, candidates[j])
			dfs(j+1, sum+candidates[j], subset, depth+1)
			subset = subset[:len(subset)-1]
		}

	}

	dfs(0, 0, nil, 0)

	return result

}

// @lc code=end
