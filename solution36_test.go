package leetcode

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValidSudoku_row(board, i) || !isValidSudoku_column(board, i) {
			return false
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !isValidSudoku_block(board, i, j) {
				return false
			}
		}
	}
	return true
}

func isValidSudoku_row(board [][]byte, row int) bool {
	key := make([]byte, 10)
	for _, v := range board[row] {
		if v == '.' {
			continue
		}
		v = v - '0'
		if v <= 0 || v > 9 || key[v] != 0 {
			return false
		}
		key[v] = v
	}
	return true
}

func isValidSudoku_column(board [][]byte, column int) bool {
	key := make([]byte, 10)
	for _, r := range board {
		v := r[column]
		if v == '.' {
			continue
		}
		v = v - '0'
		if v <= 0 || v > 9 || key[v] != 0 {
			return false
		}
		key[v] = v
	}
	return true
}

func isValidSudoku_block(board [][]byte, row int, column int) bool {
	key := make([]byte, 10)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			v := board[row+i][column+j]
			if v == '.' {
				continue
			}
			v = v - '0'
			if v <= 0 || v > 9 || key[v] != 0 {
				return false
			}
			key[v] = v
		}
	}
	return true
}
