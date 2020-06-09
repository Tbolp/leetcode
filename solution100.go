package leetcode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil {
		return false
	} else if q == nil {
		return false
	}
	if p.Val == q.Val {
		if isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) {
			return true
		}
	}
	return false
}
