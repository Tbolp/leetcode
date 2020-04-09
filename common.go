package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func createListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	start := ListNode{}
	cur := &start
	for _, val := range vals {
		cur.Next = &ListNode{}
		cur = cur.Next
		cur.Val = val
	}
	return start.Next
}
