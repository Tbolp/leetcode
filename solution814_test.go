package leetcode

func pruneTree(root *TreeNode) *TreeNode {
	parent := &TreeNode{
		Left: root,
	}
	pruneTree_(root, parent)
	return parent.Left
}

func pruneTree_(root *TreeNode, parent *TreeNode) {
	if root.Left != nil {
		pruneTree_(root.Left, root)
	}
	if root.Right != nil {
		pruneTree_(root.Right, root)
	}
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		if parent.Left == root {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
	}
}
