package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func calculateMinimumHP_old(dungeon [][]int) int {
	path := [][]int{}
	for i := 0; i < len(dungeon); i++ {
		path = append(path, make([]int, len(dungeon[0])))
	}
	minhealth := -1
	choose(0, 0, path, dungeon, 0, &minhealth, 0)
	return minhealth
}

func choose(x, y int, path, dungeon [][]int, curhealth int, minhealth *int, minpath int) {
	curhealth += dungeon[x][y]
	if *minhealth >= 0 && -curhealth > *minhealth {
		return
	}
	if curhealth < minpath {
		minpath = curhealth
	}
	path[x][y] = 1
	if x == len(dungeon)-1 && y == len(dungeon[0])-1 {
		if minpath >= 0 {
			*minhealth = 1
		} else {
			if -minpath < *minhealth || *minhealth == -1 {
				*minhealth = -minpath + 1
			}
		}
		path[x][y] = 0
		return
	}
	if x < len(dungeon)-1 && path[x+1][y] != 1 {
		choose(x+1, y, path, dungeon, curhealth, minhealth, minpath)
	}
	if y < len(dungeon[0])-1 && path[x][y+1] != 1 {
		choose(x, y+1, path, dungeon, curhealth, minhealth, minpath)
	}
	path[x][y] = 0
}

func calculateMinimumHP(dungeon [][]int) int {
	dp := [][]int{}
	for i := 0; i < len(dungeon); i++ {
		dp = append(dp, make([]int, len(dungeon[0])))
	}
	return getMin(0, 0, dp, dungeon)
}

func getMin(x, y int, dp [][]int, dungeon [][]int) int {
	if dp[x][y] != 0 {
		return dp[x][y]
	}
	if x == len(dungeon)-1 && y == len(dungeon[0])-1 {
		dp[x][y] = int(math.Max(float64(1), float64(-dungeon[x][y]+1)))
	} else if x < len(dungeon)-1 && y < len(dungeon[0])-1 {
		dp[x][y] = int(math.Max(1, float64(int(math.Min(float64(getMin(x+1, y, dp, dungeon)), float64(getMin(x, y+1, dp, dungeon))))-dungeon[x][y])))
	} else if x < len(dungeon)-1 {
		dp[x][y] = int(math.Max(1, float64(getMin(x+1, y, dp, dungeon)-dungeon[x][y])))
	} else {
		dp[x][y] = int(math.Max(1, float64(getMin(x, y+1, dp, dungeon)-dungeon[x][y])))
	}
	return dp[x][y]
}

func Test_calculateMinimumHP_Usage1(t *testing.T) {
	dungeon := createTwoDimArray([]int{-2, -3, 3, -5, -10, 1, 10, 30, -5}, 3, 3)
	if calculateMinimumHP(dungeon) != 7 {
		t.Fail()
	}
}

func Test_calculateMinimumHP_Usage2(t *testing.T) {
	dungeon := createTwoDimArray([]int{0, -74, -47, -20, -23, -39, -48, 37, -30, 37, -65, -82, 28, -27, -76, -33, 7, 42, 3, 49, -93, 37, -41, 35, -16, -96, -56, 38, -52, 19, -37, 14, -65, -42, 9, 5, -26, -30, -65, 11, 5, 16, -60, 9, 36, -36, 41, -47, -86, -22, 19, -5, -41, -8, -96, -95}, 8, 7)
	fmt.Println(calculateMinimumHP(dungeon))
}
