package leetcode

func solveSudoku(board [][]byte) {
	for {
		ninfo := false
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				for k := '1'; k < '9'+1; k++ {
					ninfo = ninfo || solveSudoku_square(board, byte(k), i*3, j*3)
				}
			}
		}
		if !ninfo {
			break
		}
	}
	solveSudoku_(board)
}

func solveSudoku_square(board [][]byte, target byte, row, col int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[row+i][col+j] == target {
				return false
			}
		}
	}
	flag := false
	r := 0
	c := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[row+i][col+j] == '.' {
				if solveSudoku_check_col(board, target, row+i) && solveSudoku_check_row(board, target, col+j) {
					if !flag {
						flag = true
						r = row + i
						c = col + j
					} else {
						return false
					}
				}
			}
		}
	}
	board[r][c] = target
	return true
}

func solveSudoku_(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for k := byte('1'); k < '9'+1; k++ {
					if solveSudoku_check_square(board, k, i, j) && solveSudoku_check_col(board, k, i) && solveSudoku_check_row(board, k, j) {
						board[i][j] = k
						if solveSudoku_(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func solveSudoku_check_square(board [][]byte, target byte, row, col int) bool {
	r := row / 3
	c := col / 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if r*3+i == row && c*3+j == col {
				continue
			}
			if board[r*3+i][c*3+j] == target {
				return false
			}
		}
	}
	return true
}

func solveSudoku_check_row(board [][]byte, target byte, col int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == target {
			return false
		}
	}
	return true
}

func solveSudoku_check_col(board [][]byte, target byte, row int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == target {
			return false
		}
	}
	return true
}
