package leetcode

func findTarget(root *TreeNode, k int) bool {
	return findTarget_(root, root, k)
}

func findTarget_(root *TreeNode, cur *TreeNode, k int) bool {
	if cur == nil {
		return false
	}
	if findTarget__(root, cur, k) {
		return true
	}
	if findTarget_(root, cur.Left, k) {
		return true
	}
	if findTarget_(root, cur.Right, k) {
		return true
	}
	return false
}

func findTarget__(root *TreeNode, node *TreeNode, k int) bool {
	cur := root
	for {
		tmp := node.Val + cur.Val
		if tmp == k {
			if node == cur {
				return false
			} else {
				return true
			}
		} else if tmp < k {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
		if cur == nil {
			return false
		}
	}
}
