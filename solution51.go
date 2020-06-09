package leetcode

func solveNQueens(n int) [][]string {
	res := [][]int{}
	placeQuene(0, n, make([]int, n), &res)
	resstr := [][]string{}
	for i := 0; i < len(res); i++ {
		resstr = append(resstr, []string{})
		for j := 0; j < n; j++ {
			tmp := ""
			for k := 0; k < n; k++ {
				if res[i][j] == k {
					tmp += "Q"
				} else {
					tmp += "."
				}
			}
			resstr[i] = append(resstr[i], tmp)
		}
	}
	return resstr
}

func placeQuene(i int, n int, tmp []int, res *[][]int) {
	if i == n {
		b := make([]int, len(tmp))
		copy(b, tmp)
		*res = append(*res, b)
		return
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
	for ind, candidate := range candidates {
		if candidate == false {
			continue
		}
		tmp[i] = ind
		placeQuene(i+1, n, tmp, res)
	}
}
