package leetcode

func goodNodes(root *TreeNode) int {
	queue := []*TreeNode{root}
	queue_max := []int{-100000}
	sum := 0
	for len(queue) != 0 {
		cur := queue[0]
		max := queue_max[0]
		if max <= cur.Val {
			sum++
		}
		if cur.Left != nil {
			queue = append(queue, cur.Left)
			if cur.Val > max {
				queue_max = append(queue_max, cur.Val)
			} else {
				queue_max = append(queue_max, max)
			}
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
			if cur.Val > max {
				queue_max = append(queue_max, cur.Val)
			} else {
				queue_max = append(queue_max, max)
			}
		}
		queue = queue[1:]
		queue_max = queue_max[1:]
	}
	return sum
}
