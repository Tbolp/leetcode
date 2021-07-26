package leetcode

import "testing"

func sortList(head *ListNode) *ListNode {
	seglen := 2
	current := head
	totallen := 0
	for current != nil {
		totallen++
		current = current.Next
	}
	tmp := &ListNode{Next: head}
	head = tmp
	first := head
	second := head
	third := head
	for seglen < totallen*2 {
		first = head
		for {
			second = sortList_Move(first, seglen/2)
			if second == nil || second.Next == nil {
				break
			}
			third = sortList_Move(second, seglen/2)
			if third != nil {
				third = third.Next
			}
			first = sortList_Merge(first, second, third)
		}
		seglen = seglen * 2
	}
	return head.Next
}

func sortList_Move(head *ListNode, count int) *ListNode {
	second := head
	for i := 0; i < count; i++ {
		second = second.Next
		if second == nil {
			return nil
		}
	}
	return second
}

func sortList_Merge(head *ListNode, inner *ListNode, tail *ListNode) *ListNode {
	first := head.Next
	second := inner.Next
	third := tail
	inner.Next = nil
	current := head
	for first != nil || second != third {
		if first != nil && second != third {
			if first.Val < second.Val {
				current.Next = first
				first = first.Next
			} else {
				current.Next = second
				second = second.Next
			}
		} else if first != nil {
			current.Next = first
			first = first.Next
		} else {
			current.Next = second
			second = second.Next
		}
		current = current.Next
	}
	current.Next = third
	return current
}

func Test_sortList_Usage1(t *testing.T) {
	sortList(createListNode([]int{4, 2, 3, 1}))
}
