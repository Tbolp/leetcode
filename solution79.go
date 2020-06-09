package leetcode

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				path := make([][]bool, len(board))
				for k := 0; k < len(path); k++ {
					path[k] = make([]bool, len(board[k]))
				}
				if _searchWord(board, word, i, j, path) {
					return true
				}
			}
		}
	}
	return false
}

func _searchWord(board [][]byte, word string, i int, j int, path [][]bool) bool {
	if len(word) == 1 {
		return true
	}
	path[i][j] = true
	if i-1 >= 0 && path[i-1][j] == false && board[i-1][j] == word[1] {
		if _searchWord(board, word[1:], i-1, j, path) {
			return true
		}
	}
	if i+1 < len(path) && path[i+1][j] == false && board[i+1][j] == word[1] {
		if _searchWord(board, word[1:], i+1, j, path) {
			return true
		}
	}
	if j-1 >= 0 && path[i][j-1] == false && board[i][j-1] == word[1] {
		if _searchWord(board, word[1:], i, j-1, path) {
			return true
		}
	}
	if j+1 < len(path[i]) && path[i][j+1] == false && board[i][j+1] == word[1] {
		if _searchWord(board, word[1:], i, j+1, path) {
			return true
		}
	}
	path[i][j] = false
	return false
}
