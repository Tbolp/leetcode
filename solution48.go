package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			temp := matrix[i][j]
			ti := i
			tj := j
			for k := 0; k < 3; k++ {
				matrix[ti][tj] = matrix[n-1-tj][ti]
				nti := n - 1 - tj
				ntj := ti
				ti = nti
				tj = ntj
			}
			matrix[ti][tj] = temp
		}
	}
}
