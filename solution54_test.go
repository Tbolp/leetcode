package leetcode

func spiralOrder(matrix [][]int) []int {
	right := len(matrix[0])
	down := len(matrix)
	up := -1
	left := -1
	x := 0
	y := 0
	res := make([]int, right*down)
	i := 0
	const r = 0
	const d = 1
	const u = 2
	const l = 4
	status := r
	for i != len(res) {
		res[i] = matrix[x][y]
		i++
		switch status {
		case r:
			if y+1 < right {
				y++
			} else {
				status = d
				x++
				up++
			}
		case d:
			if x+1 < down {
				x++
			} else {
				status = l
				y--
				right--
			}
		case l:
			if y-1 > left {
				y--
			} else {
				status = u
				x--
				down--
			}
		case u:
			if x-1 > up {
				x--
			} else {
				status = r
				y++
				left++
			}
		}
	}
	return res
}
