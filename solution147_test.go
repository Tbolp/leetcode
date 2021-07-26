package leetcode

func insertionSortList(head *ListNode) *ListNode {
	output := &ListNode{}
	for head != nil {
		next := head.Next
		if output.Next == nil {
			output.Next = head
			head.Next = nil
		} else {
			pre := output
			cur := output.Next
			flag := false
			for cur != nil {
				if head.Val < cur.Val {
					pre.Next = head
					head.Next = cur
					flag = true
					break
				}
				pre = cur
				cur = cur.Next
			}
			if !flag {
				pre.Next = head
				head.Next = nil
			}
		}
		head = next
	}
	return output.Next
}
