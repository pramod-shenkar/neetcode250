package neetcode250

import "slices"

/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start

func fourSum(nums []int, target int) [][]int {

	var result = [][]int{}

	slices.Sort(nums)

	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			var start, end = j + 1, len(nums) - 1

			for start < end {

				got := nums[i] + nums[j] + nums[start] + nums[end]
				if got == target {
					result = append(result, []int{nums[i], nums[j], nums[start], nums[end]})

					for start < end && nums[start] == nums[start+1] {
						start++
					}

					for start < end && nums[end] == nums[end-1] {
						end--
					}

					start++
					end--
				} else if got < target {
					start++
				} else if got > target {
					end--
				}
			}

		}
	}

	return result
}

// @lc code=end
