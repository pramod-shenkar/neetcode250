package neetcode250

import (
	"strings"
)

/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {

	/*
		"/home/user/Documents/../Pictures"
	*/

	var s []string

	r := strings.Split(path, "/")
	for _, v := range r {

		if v == "" || v == "." {
			continue
		} else if v == ".." {
			if len(s) > 0 {
				s = s[0 : len(s)-1]
			}
		} else {
			s = append(s, v)
		}

	}

	return "/" + strings.Join(s, "/")
}

// @lc code=end

/*
- split string now if . or "" then continue,
if ".." then top from stack
if anything else ->  just add to stack

*/
