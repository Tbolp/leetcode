package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	fakehead := &ListNode{Next: head}
	cur := fakehead
	for {
		next := cur
		for i := 0; i < k; i++ {
			if next == nil {
				break
			}
			next = next.Next
		}
		if cur == nil || next == nil {
			break
		}
		next = cur.Next
		for i := k - 1; i > 0; i-- {
			reverseKGroup_move(cur, i)
		}
		cur = next
	}
	return fakehead.Next
}

func reverseKGroup_move(head *ListNode, k int) {
	pre := head
	for i := 0; i < k; i++ {
		cur := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = cur
		pre = pre.Next
	}
}
