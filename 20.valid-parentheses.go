package neetcode250

import "fmt"

/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	var stack []rune

	for _, v := range s {
		switch v {
		case '{', '(', '[':
			stack = append(stack, v)
		default:
			if len(stack) == 0 {
				return false
			}
			got := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if got == '{' && v == '}' ||
				got == '(' && v == ')' ||
				got ==
					'[' && v == ']' {
				continue
			} else {
				fmt.Println(string(got), string(v))
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

// @lc code=end

/*

things to remeber :
- Never forgot edge cases ie "{", "}{" "{{"

*/
