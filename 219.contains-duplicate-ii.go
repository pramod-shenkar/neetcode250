package neetcode250

/*
 * @lc app=leetcode id=219 lang=golang
 *
 * [219] Contains Duplicate II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {

	var m = map[int]int{}

	for i, v := range nums {

		if j, isExist := m[v]; isExist {
			if i-j <= k {
				return true
			}
		}

		m[v] = i
	}

	return false
}

// @lc code=end
