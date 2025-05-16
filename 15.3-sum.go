package neetcode250

import (
	"slices"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {

	slices.Sort(nums)

	var result = [][]int{}

	for i := range nums {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		var start, end = i + 1, len(nums) - 1
		for start < end {

			var target = nums[i] + nums[start] + nums[end]

			if target == 0 {
				result = append(result, []int{nums[i], nums[start], nums[end]})

				for start > len(nums)-1 && nums[start] == nums[start+1] {
					start++
				}
				for end > 1 && nums[end] == nums[end-1] {
					end--
				}

				start++
				end--

			} else if target < 0 {
				start++
			} else {
				end--
			}

		}

	}

	return result

}

// @lc code=end
