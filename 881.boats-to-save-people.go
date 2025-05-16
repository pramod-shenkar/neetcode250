package neetcode250

import (
	"slices"
)

/*
 * @lc app=leetcode id=881 lang=golang
 *
 * [881] Boats to Save People
 */

// @lc code=start
func numRescueBoats(people []int, limit int) int {

	slices.Sort(people)

	var result int

	var left, right = 0, len(people) - 1

	for left <= right { // needing <= cause at last one person can remained
		if people[left]+people[right] <= limit {
			left++
			right--
		} else {
			right--
		}
		result++

	}

	return result
}

// @lc code=end

/* 	1	2	2	3



 */
