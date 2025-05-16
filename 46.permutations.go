package neetcode250

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	var result [][]int

	var backtrack func(permuation []int, visited []bool)
	backtrack = func(permuation []int, visited []bool) {

		if len(permuation) == len(nums) {
			var tempset = make([]int, len(permuation))
			copy(tempset, permuation)
			result = append(result, tempset)
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}

			permuation = append(permuation, nums[i])
			visited[i] = true
			backtrack(permuation, visited)
			visited[i] = false

			permuation = permuation[:len(permuation)-1]
		}
	}

	var visited = make([]bool, len(nums))
	backtrack(nil, visited)

	return result
}

// @lc code=end
