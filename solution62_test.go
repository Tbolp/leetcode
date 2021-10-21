package leetcode

func uniquePaths(m int, n int) int {
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		cache[i][0] = 1
	}
	for j := 0; j < n; j++ {
		cache[0][j] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if cache[i][j] != 0 {
				continue
			} else {
				cache[i][j] = cache[i-1][j] + cache[i][j-1]
			}
		}
	}
	return cache[m-1][n-1]
}
