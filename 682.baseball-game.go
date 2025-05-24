package neetcode250

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=682 lang=golang
 *
 * [682] Baseball Game
 */

// @lc code=start
func calPoints(operations []string) int {

	// stack -----------------------

	type stack []int
	var push = func(s *stack, i int) {
		*s = append(*s, i)
	}

	var get = func(s *stack, i int) int {
		if len(*s)-i > 0 {
			last := (*s)[len(*s)-1-i]
			return last
		}
		return 0

	}
	var pop = func(s *stack) int {
		if len(*s) > 0 {
			last := (*s)[len(*s)-1]
			*s = (*s)[:len(*s)-1]
			return last
		}
		return 0
	}

	_, _ = push, pop

	// impl -----------------------

	var s = &stack{}

	for _, op := range operations {
		switch op {
		case "+":
			num1 := get(s, 0)
			num2 := get(s, 1)

			push(s, num1+num2)

		case "D":
			num1 := get(s, 0)
			push(s, num1*2)

		case "C":
			pop(s)

		default:
			score, _ := strconv.Atoi(op)
			push(s, score)
		}

		fmt.Println(s)

	}

	var result int
	for _, v := range *s {
		result += v
	}

	return result

}

// @lc code=end
