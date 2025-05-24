package neetcode250

/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {

	var result []int = make([]int, len(temperatures))

	var s []struct {
		index int
		temp  int
	}

	for i, v := range temperatures {

		if len(s) > 0 {
			for len(s) > 0 && s[len(s)-1].temp < v {
				result[s[len(s)-1].index] = i - s[len(s)-1].index
				s = s[:len(s)-1]
			}
		}

		s = append(s, struct {
			index int
			temp  int
		}{i, v})

	}

	return result
}

// @lc code=end


/* 
note:
each time you have to add temp in stack. Just condtion is if
	- if stack empty so noone there to compare :  add to stack
	- if stack not empty so we can compare
		- check current temp is more than stack.top, 
			- mark stack.top in result 
			- remove stack.top its no longer needed & result is marked
			do this until u get smaller temp or stack is empty
		- now current to stakc
*/
