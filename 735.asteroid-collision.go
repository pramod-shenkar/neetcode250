package neetcode250

import "fmt"

/*
 * @lc app=leetcode id=735 lang=golang
 *
 * [735] Asteroid Collision
 */

// @lc code=start
func asteroidCollision(asteroids []int) []int {

	var s []int

	for i := 0; i < len(asteroids); i++ {

		num1 := asteroids[i]
		if len(s) == 0 {
			s = append(s, num1)
		} else {
			num2 := s[len(s)-1]
			if num1*num2 < 0 {
				s = s[:len(s)-1]
				if num1+num2 != 0 {
					s = append(s, MaxSign(num1, num2))
				}
			} else if num1*num2 > 0 {
				s = append(s, num1)
			}
		}
		fmt.Println(s)
	}

	var colided bool = true
	for colided {
		colided = false

		if len(s) <= 1 {
			break
		}

		num1 := s[len(s)-1]
		num2 := s[len(s)-2]

		if diffSign(num1, num2) {
			s = s[:len(s)-2]
			s = append(s, MaxSign(num1, num2))
			colided = true
		}

	}

	return s
}

func diffSign(i, j int) bool {
	return i*j <= 0
}

func MaxSign(i, j int) int {
	var a, b int
	if i < 0 {
		a = i * -1
	} else {
		a = j
	}

	if j < 0 {
		b = j * -1
	} else {
		b = j
	}

	if a < b {
		return j
	} else {
		return i
	}
}

// @lc code=end
