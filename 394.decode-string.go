package neetcode250

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func decodeString(str string) string {

	var s []string

	for _, v := range str {

		if v == ']' {

			var temp string
			for s[len(s)-1] != "[" {
				temp = string(s[len(s)-1]) + temp
				s = s[:len(s)-1]
			}
			s = s[:len(s)-1]

			var count int
			var denom = 1
			for len(s) > 0 {
				countStr := s[len(s)-1]
				c, err := strconv.Atoi(string(countStr))
				if err != nil {
					break
				}
				count += c * denom
				denom = denom * 10
				s = s[:len(s)-1]
			}

			fmt.Println(count)

			// s = s[:len(s)-1]

			// fmt.Println("adding", temp, count, strings.Repeat(temp, count))

			s = append(s, strings.Repeat(temp, count))

		} else {
			s = append(s, string(v))
		}

		// fmt.Println(string(v), s)
	}

	return strings.Join(s, "")
}

// @lc code=end
