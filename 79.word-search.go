package neetcode250

/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {

	var result bool
	var dfs func(i int, j int, wordIndex int, depth int)
	dfs = func(i int, j int, wordIndex int, depth int) {

		if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
			return
		}

		if word[wordIndex] != board[i][j] {
			return
		}

		if wordIndex == len(word)-1 {
			result = true
			return
		}

		// log.Println(strings.Repeat("    ", depth), string(word[wordIndex]))

		var temp byte = '0'
		board[i][j], temp = temp, board[i][j]
		dfs(i+1, j, wordIndex+1, depth+1)
		dfs(i-1, j, wordIndex+1, depth+1)
		dfs(i, j+1, wordIndex+1, depth+1)
		dfs(i, j-1, wordIndex+1, depth+1)
		board[i][j], temp = temp, board[i][j]

	}

	for i, row := range board {
		for j, col := range row {
			if col == word[0] {
				dfs(i, j, 0, 0)
			}
		}
	}
	return result
}

// @lc code=end
