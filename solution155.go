package leetcode

import "math"

type element struct {
	elt    int
	minElt int
	prev   *element
}

type MinStack struct {
	top *element
}

func (this *MinStack) Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if this.top == nil {
		this.top = &element{elt: x, prev: nil, minElt: x}
	} else {
		elt := element{elt: x, prev: this.top, minElt: int(math.Min(float64(this.top.minElt), float64(x)))}
		this.top = &elt
	}
}

func (this *MinStack) Pop() {
	if this.top != nil {
		this.top = this.top.prev
	}
}

func (this *MinStack) Top() int {
	return this.top.elt
}

func (this *MinStack) GetMin() int {
	return this.top.minElt
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
