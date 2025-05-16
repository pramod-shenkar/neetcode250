package neetcode250

/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {

	var start, end = 0, len(numbers)-1

	for start <= end {
		if numbers[start]+numbers[end] == target {
			return []int{start+1, end+1}
		} else {

			if numbers[start]+numbers[end] < target {
				start++
			} else {
				end--
			}

		}

	}

	return nil
}

// @lc code=end
