package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	k := -1
	for i := 0; i < m; i++ {
		if target >= matrix[i][0] && target <= matrix[i][n-1] {
			k = i
		}
	}
	if k == -1 {
		return false
	}
	s := 0
	e := n
	for s != e {
		m := (s + e) / 2
		val := matrix[k][m]
		if val == target {
			return true
		} else if target > val {
			s = m + 1
		} else {
			e = m
		}
	}
	return false
}
