package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
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

func createTwoDimArray(data []int, m, n int) [][]int {
	ret := [][]int{}
	for i := 0; i < m; i++ {
		ret = append(ret, make([]int, n))
		for j := 0; j < n; j++ {
			ret[i][j] = data[i*n+j]
		}
	}
	return ret
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
