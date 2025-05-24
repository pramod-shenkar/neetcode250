package neetcode250

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	stack   []int
	minimum []int
}

// func Constructor() MinStack {
// 	return MinStack{}
// }

func (this *MinStack) Push(val int) {
	if this.GetMin() == -9999 {
		this.minimum = append(this.minimum, val)
	} else {
		this.minimum = append(this.minimum, min(this.GetMin(), val))
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.minimum = this.minimum[:len(this.minimum)-1]
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	// fmt.Println(this.stack, "\n", this.minimum)
	if len(this.minimum) == 0 {
		return -9999
	}
	return this.minimum[len(this.minimum)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

/* note:
while push, we setting min by checkig existing min
but what if min is not exist? then set val as first min
*/
