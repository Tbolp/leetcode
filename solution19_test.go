package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	len := 0
	current := head
	for current != nil {
		len++
		current = current.Next
	}
	if len == 0 {
		return head
	}
	ind := len - n
	if ind < 0 {
		return head
	}
	if ind == 0 {
		return head.Next
	} else if ind == len-1 {
		pre := head
		pre = nil
		current = head
		for current.Next != nil {
			pre = current
			current = current.Next
		}
		if pre == nil {
			return nil
		} else {
			pre.Next = nil
			return head
		}
	} else {
		pre := head
		pre = nil
		next := head.Next
		current = head
		for i := 0; i < ind; i++ {
			pre = current
			current = current.Next
			next = current.Next
		}
		pre.Next = next
		return head
	}
}
