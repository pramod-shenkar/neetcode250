package neetcode250

/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {

	var result []string

	var dfs func(str string, open, close int)
	dfs = func(str string, open, close int) {

		if close > open {
			return
		}

		// breaking condition :  if length maches
		if open == close && len(str) == n*2 {
			result = append(result, str)
			return
		}

		if open < n {
			dfs(str+"(", open+1, close)
		}

		if close < n {
			dfs(str+")", open, close+1)
		}

	}

	dfs("(", 1, 0)

	return result
}

// @lc code=end

/*

- use backtracking:
	return condition  => if closing brackets are more than open brackets
	append if => substring get completed ie len=n*2
	backtrack for
		opening/closing bracket not yet reached to n

	(
	/\
   (  )
   /\
  (  )


*/
