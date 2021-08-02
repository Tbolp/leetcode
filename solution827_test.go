package leetcode

func largestIsland(grid [][]int) int {
	largest := 0
	color := 2
	areas := map[int]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				areas[color] = largestIsland_color(grid, color, i, j)
				color++
			}
		}
	}
	for _, v := range areas {
		if v > largest {
			largest = v
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				tmp := largestIsland_calc(grid, i, j, areas)
				if tmp > largest {
					largest = tmp
				}
			}
		}
	}
	return largest
}

func largestIsland_color(grid [][]int, color int, i int, j int) int {
	sum := 1
	grid[i][j] = color
	if i > 0 && grid[i-1][j] == 1 {
		sum += largestIsland_color(grid, color, i-1, j)
	}
	if j > 0 && grid[i][j-1] == 1 {
		sum += largestIsland_color(grid, color, i, j-1)
	}
	if i < len(grid)-1 && grid[i+1][j] == 1 {
		sum += largestIsland_color(grid, color, i+1, j)
	}
	if j < len(grid[0])-1 && grid[i][j+1] == 1 {
		sum += largestIsland_color(grid, color, i, j+1)
	}
	return sum
}

func largestIsland_calc(grid [][]int, i int, j int, areas map[int]int) int {
	adj := map[int]bool{}
	if i > 0 {
		adj[grid[i-1][j]] = true
	}
	if j > 0 {
		adj[grid[i][j-1]] = true
	}
	if i < len(grid)-1 {
		adj[grid[i+1][j]] = true
	}
	if j < len(grid[0])-1 {
		adj[grid[i][j+1]] = true
	}
	sum := 1
	for k := range adj {
		sum += areas[k]
	}
	return sum
}
