package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	status := 1
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			status = 0
		}
		cache[i][0] = status
	}
	status = 1
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			status = 0
		}
		cache[0][j] = status
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if cache[i][j] != 0 {
				continue
			} else {
				cache[i][j] = cache[i-1][j] + cache[i][j-1]
			}
		}
	}
	return cache[m-1][n-1]
}
