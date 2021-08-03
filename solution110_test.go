package leetcode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ll := 0
	if root.Left != nil {
		if !isBalanced(root.Left) {
			return false
		}
		ll = root.Left.Val
	}
	lr := 0
	if root.Right != nil {
		if !isBalanced(root.Right) {
			return false
		}
		lr = root.Right.Val
	}
	detal := ll - lr
	if detal == 0 || detal == 1 || detal == -1 {
		if lr > ll {
			root.Val = lr + 1
		} else {
			root.Val = ll + 1
		}
		return true
	} else {
		return false
	}
}
