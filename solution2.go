package leetcode

func calc(flag bool, v1 int, v2 int) (bool, int) {
	var v int
	if flag {
		v = v1 + v2 + 1
	} else {
		v = v1 + v2
	}
	if v/10 != 0 {
		flag = true
	} else {
		flag = false
	}
	v = v % 10
	return flag, v
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	node1 := l1
	node2 := l2
	startnode := ListNode{}
	node3 := &startnode
	flag := false
	for node1 != nil || node2 != nil {
		node3.Next = &ListNode{}
		node3 = node3.Next
		if node1 == nil {
			flag, node3.Val = calc(flag, 0, node2.Val)
			node2 = node2.Next
		} else if node2 == nil {
			flag, node3.Val = calc(flag, node1.Val, 0)
			node1 = node1.Next
		} else {
			flag, node3.Val = calc(flag, node1.Val, node2.Val)
			node1 = node1.Next
			node2 = node2.Next
		}
	}
	if flag {
		node3.Next = &ListNode{}
		node3 = node3.Next
		node3.Val = 1
	}
	return startnode.Next
}
