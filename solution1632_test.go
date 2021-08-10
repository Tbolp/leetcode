package leetcode

var matrixRankTransform_Context struct {
	Matrix [][]int
	Output [][]int
}

func matrixRankTransform(matrix [][]int) [][]int {
	matrixRankTransform_Context.Matrix = matrix
	matrixRankTransform_Context.Output = make([][]int, len(matrix))
	for i := range matrixRankTransform_Context.Matrix {
		matrixRankTransform_Context.Output[i] = make([]int, len(matrix[i]))
	}
	for i := range matrix {
		for j := range matrix[0] {
			val := matrixRankTransform_getInit(i, j)
			matrixRankTransform_change(i, j, val)
		}
	}
	return matrixRankTransform_Context.Output
}

func matrixRankTransform_getInit(row int, col int) int {
	val := 1
	for i := 0; i < len(matrixRankTransform_Context.Matrix); i++ {
		if i == row {
			continue
		}
		if matrixRankTransform_Context.Output[i][col] == 0 {
			continue
		}
		if matrixRankTransform_Context.Matrix[i][col] < matrixRankTransform_Context.Matrix[row][col] {
			if matrixRankTransform_Context.Output[i][col]+1 > val {
				val = matrixRankTransform_Context.Output[i][col] + 1
			}
		} else if matrixRankTransform_Context.Matrix[i][col] == matrixRankTransform_Context.Matrix[row][col] {
			if matrixRankTransform_Context.Output[i][col] > val {
				val = matrixRankTransform_Context.Output[i][col]
			}
		}
	}

	for i := 0; i < len(matrixRankTransform_Context.Matrix[0]); i++ {
		if i == col {
			continue
		}
		if matrixRankTransform_Context.Output[row][i] == 0 {
			continue
		}
		if matrixRankTransform_Context.Matrix[row][i] < matrixRankTransform_Context.Matrix[row][col] {
			if matrixRankTransform_Context.Output[row][i]+1 > val {
				val = matrixRankTransform_Context.Output[row][i] + 1
			}
		} else if matrixRankTransform_Context.Matrix[row][i] == matrixRankTransform_Context.Matrix[row][col] {
			if matrixRankTransform_Context.Output[row][i] > val {
				val = matrixRankTransform_Context.Output[row][i]
			}
		}
	}
	return val
}

func matrixRankTransform_change(row int, col int, val int) {
	matrixRankTransform_Context.Output[row][col] = val
	for i := 0; i < len(matrixRankTransform_Context.Matrix); i++ {
		if matrixRankTransform_Context.Output[i][col] == 0 {
			continue
		}
		if matrixRankTransform_Context.Matrix[i][col] > matrixRankTransform_Context.Matrix[row][col] {
			if matrixRankTransform_Context.Output[i][col] <= matrixRankTransform_Context.Output[row][col] {
				matrixRankTransform_change(i, col, matrixRankTransform_Context.Output[row][col]+1)
			}
		} else if matrixRankTransform_Context.Matrix[i][col] == matrixRankTransform_Context.Matrix[row][col] {
			if matrixRankTransform_Context.Output[i][col] != matrixRankTransform_Context.Output[row][col] {
				matrixRankTransform_change(i, col, matrixRankTransform_Context.Output[row][col])
			}
		}
	}

	for i := 0; i < len(matrixRankTransform_Context.Matrix[0]); i++ {
		if matrixRankTransform_Context.Output[row][i] == 0 {
			continue
		}
		if matrixRankTransform_Context.Matrix[row][i] > matrixRankTransform_Context.Matrix[row][col] {
			if matrixRankTransform_Context.Output[row][i] <= matrixRankTransform_Context.Output[row][col] {
				matrixRankTransform_change(row, i, matrixRankTransform_Context.Output[row][col]+1)
			}
		} else if matrixRankTransform_Context.Matrix[row][i] == matrixRankTransform_Context.Matrix[row][col] {
			if matrixRankTransform_Context.Output[row][i] != matrixRankTransform_Context.Output[row][col] {
				matrixRankTransform_change(row, i, matrixRankTransform_Context.Output[row][col])
			}
		}
	}

}
