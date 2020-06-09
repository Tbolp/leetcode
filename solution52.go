package leetcode

func totalNQueens(n int) int {
	return _totalNQuenes(0, n, make([]int, n))
}

func _totalNQuenes(i int, n int, tmp []int) int {
	if i == n {
		return 1
	}
	candidates := make([]bool, n)
	for j := 0; j < n; j++ {
		candidates[j] = true
	}
	for j := 0; j < i; j++ {
		candidates[tmp[j]] = false
		if tmp[j]+i-j < n {
			candidates[tmp[j]+i-j] = false
		}
		if tmp[j]-i+j >= 0 {
			candidates[tmp[j]-i+j] = false
		}
	}
	sum := 0
	for ind, candidate := range candidates {
		if candidate == false {
			continue
		}
		tmp[i] = ind
		sum += _totalNQuenes(i+1, n, tmp)
	}
	return sum
}
