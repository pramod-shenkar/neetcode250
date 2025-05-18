package neetcode250

/*
 * @lc app=leetcode id=2461 lang=golang
 *
 * [2461] Maximum Sum of Distinct Subarrays With Length K
 */

// @lc code=start
func maximumSubarraySum(nums []int, k int) int64 {

	var result int

	var sum int

	var m = map[int]int{}

	for i := 0; i < k; i++ {
		sum += nums[i]
		m[i]++
	}


	result = sum

	for i := 0; i < len(nums)-k; i++ {
		if nums[i+k] == nums[i+k-1] {
			continue
		}
		sum = sum - nums[i] + nums[i+k]
		result = max(result, sum)
	}

	return int64(result)

}

// @lc code=end
