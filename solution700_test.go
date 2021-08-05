package leetcode

func searchBST(root *TreeNode, val int) *TreeNode {
	cur := root
	for cur != nil {
		if cur.Val == val {
			return cur
		}
		if val < cur.Val {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return cur
}
