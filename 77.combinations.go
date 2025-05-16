package neetcode250

/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {

	var result [][]int

	var dfs func(i int, subset []int, depth int)
	dfs = func(i int, subset []int, depth int) {

		if len(subset) > k {
			return
		}

		// log.Printf("%s %v ", strings.Repeat("    ", depth), subset)

		if len(subset) == k {
			var tempset = make([]int, len(subset))
			copy(tempset, subset)
			result = append(result, tempset)
			return
		}

		for j := i; j <= n; j++ {
			subset = append(subset, j)
			dfs(j+1, subset, depth+1)
			subset = subset[:len(subset)-1]
		}

	}

	dfs(1, nil, 0)

	return result

}

// @lc code=end
