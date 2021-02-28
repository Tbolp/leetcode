package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	var current *ListNode
	var min *ListNode
	for {
		min = nil
		for _, v := range lists {
			if v != nil {
				if min == nil {
					min = v
				} else if v.Val < min.Val {
					min = v
				}
			}
		}
		if min == nil {
			break
		}
		for i, v := range lists {
			if v == min {
				lists[i] = v.Next
				break
			}
		}
		if head == nil {
			head = min
			current = min
		} else {
			current.Next = min
			current = current.Next
		}
	}
	if head != nil {
		current.Next = nil
	}
	return head
}
