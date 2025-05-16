package neetcode250

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start

func subsets(nums []int) [][]int {

	var result [][]int

	var subset stack

	var dfs func(int)
	dfs = func(i int) {

		if i > len(nums)-1 {
			var tempSet = make([]int, len(subset))
			copy(tempSet, subset)
			result = append(result, tempSet)
			return
		}

		// choose element
		subset.push(nums[i])
		dfs(i + 1)

		// not choose element
		subset.pop()
		dfs(i + 1)

	}

	dfs(0)

	return result
}

type stack []int

func (a *stack) push(i int) {
	*a = append(*a, i)
}

func (a *stack) pop() {
	*a = (*a)[:len(*a)-1]
}

// @lc code=end
