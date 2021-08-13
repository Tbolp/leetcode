package leetcode

func setZeroes(matrix [][]int) {
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	for i := range rows {
		if !rows[i] {
			continue
		}
		for j := range matrix[0] {
			matrix[i][j] = 0
		}
	}
	for i := range cols {
		if !cols[i] {
			continue
		}
		for j := range matrix {
			matrix[j][i] = 0
		}
	}
}
