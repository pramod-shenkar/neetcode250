package neetcode250

import (
	"log"
	"slices"
	"strings"
)

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {

	var result [][]int

	slices.Sort(nums)

	var dfs func(i int, permutation []int, visted []bool, depth int)
	dfs = func(i int, permutation []int, visted []bool, depth int) {

		// if len(nums) > i {
		log.Printf("%s %v%v", strings.Repeat("    ", depth), permutation, 0)
		// }

		if len(permutation) == len(nums) {
			result = append(result, slices.Clone(permutation))
			return
		}

		for j := 0; j < len(nums); j++ {

			if visted[j] {
				continue
			}

			// if its duplicate & previous same no already visited then dont use current
			if j > 0 && nums[j] == nums[j-1] && !visted[j-1] {
				continue
			}

			permutation = append(permutation, nums[j])
			visted[j] = true
			dfs(j+1, permutation, visted, depth+1)
			visted[j] = false
			permutation = permutation[:len(permutation)-1]
		}
	}

	// for i := range nums {
	var visited = make([]bool, len(nums))
	dfs(0, nil, visited, 0)
	// }

	return result
}

// @lc code=end
