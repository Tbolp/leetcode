package solution382

import "math/rand"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	data *ListNode
}

func Constructor(head *ListNode) Solution {
	obj := Solution{}
	obj.data = head
	return obj
}

func (this *Solution) GetRandom() int {
	ptr := this.data.Next
	val := this.data.Val
	i := 2
	for ptr != nil {
		if rand.Intn(i) == 0 {
			val = ptr.Val
		}
		i++
		ptr = ptr.Next
	}
	return val
}
