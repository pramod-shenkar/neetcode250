package neetcode250

/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int) {

	n := len(nums)
	if n <= 1 || k%n == 0 {
		return
	}

	k = k % n
	ans := make([]int, 0, n)

	left := 0
	right := n - k

	for len(ans) < n {
		if right < n {
			ans = append(ans, nums[right])
			right++
		} else {
			ans = append(ans, nums[left])
			left++
		}
	}

	copy(nums, ans)
}

// @lc code=end
