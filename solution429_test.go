package leetcode

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*Node{}
	queue = append(queue, root)
	next := 1
	res := [][]int{}
	tmp := []int{}
	sum := 0
	for len(queue) != 0 {
		tmp = append(tmp, queue[0].Val)
		next--
		sum += len(queue[0].Children)
		for _, v := range queue[0].Children {
			queue = append(queue, v)
		}
		if next == 0 {
			res = append(res, tmp)
			tmp = []int{}
			next = sum
			sum = 0
		}
		queue = queue[1:]
	}
	return res
}
