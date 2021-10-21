package leetcode

func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) != 0 {
		if queue[0] != nil && queue[1] != nil {
			if queue[0].Val != queue[1].Val {
				return false
			}
			queue = append(queue, queue[0].Left, queue[1].Right)
			queue = append(queue, queue[0].Right, queue[1].Left)
		} else if queue[0] != nil || queue[1] != nil {
			return false
		}
		queue = queue[2:]
	}
	return true
}
