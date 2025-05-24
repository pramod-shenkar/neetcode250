package neetcode250

/*
 * @lc app=leetcode id=901 lang=golang
 *
 * [901] Online Stock Span
 */

// @lc code=start
type StockSpanner struct {
	s []struct {
		price int
		span  int
	}
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {

	var span int = 1
	for len(this.s) > 0 && this.s[len(this.s)-1].price <= price {
		span += this.s[len(this.s)-1].span
		this.s = this.s[:len(this.s)-1]
	}

	this.s = append(this.s, struct {
		price int
		span  int
	}{price, span})

	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
// @lc code=end

/*
very similar to daily temprature question
idea :
harbaar stack me element push krna hi hai..
bas jab chhota element aye to itna clear hai ki top usase bada hai means no lesser element exist before ie 1
jab bada element aayet tab, pop elements from stack until found bigger
bas span 1 se nahi badega.. Wo jo jo element pop hua hai unse badhega

ie
elements:	80	70	60	50
spans:		2	2	3	1

now element is 75 means
span of 75 =  1(own) + 1(50) + 3(60) + 2(70) = 7
why: span of 60 denotes that previously we got 3 elements that less than 60 although those are less so we removed them
*/
