package neetcode250

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 */

// @lc code=start
func evalRPN(tokens []string) int {

	var s []int

	for _, v := range tokens {

		switch v {
		case "+", "-", "*", "/":

			if len(s) == 0 {
				break
			}

			num1 := s[len(s)-2]
			num2 := s[len(s)-1]
			s = s[:len(s)-2]

			var ops int
			if v == "+" {
				ops = num1 + num2
			} else if v == "-" {
				ops = num1 - num2
			} else if v == "*" {
				ops = num1 * num2
			} else if v == "/" {
				ops = num1 / num2
			}

			s = append(s, ops)
		default:
			num, _ := strconv.Atoi(v)
			s = append(s, num)
		}

		fmt.Println(s, v)

	}

	if len(s) != 1 {
		return 0
	}

	return s[0]

}

// @lc code=end
