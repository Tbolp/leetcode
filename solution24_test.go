package leetcode

func swapPairs(head *ListNode) *ListNode {
	fakehead := &ListNode{}
	fakehead.Next = head
	pre := fakehead
	cur := fakehead.Next
	for cur != nil && cur.Next != nil {
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = cur
		pre = cur
		cur = cur.Next
	}
	return fakehead.Next
}
