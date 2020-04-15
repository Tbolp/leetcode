package leetcode

func bstToGstEx(root *TreeNode, sum *int) *TreeNode {
	if root == nil {
		return nil
	}
	bstToGstEx(root.Right, sum)
	root.Val = root.Val + *sum
	*sum = root.Val
	bstToGstEx(root.Left, sum)
	return root
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	return bstToGstEx(root, &sum)
}
