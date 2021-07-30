package leetcode

import (
	"fmt"
	"testing"
)

func updateMatrix(mat [][]int) [][]int {
	num := 1
	for {
		flag := false
		for i := 0; i < len(mat); i++ {
			for j := 0; j < len(mat[0]); j++ {
				if mat[i][j] == num {
					if !updateMatrix_exist(mat, i, j, num-1) {
						flag = true
						mat[i][j] = num + 1
					}
				}
			}
		}
		if !flag {
			break
		}
		num++
	}
	return mat
}

func updateMatrix_exist(mat [][]int, i int, j int, c int) bool {
	if i > 0 {
		if mat[i-1][j] == c {
			return true
		}
	}
	if j > 0 {
		if mat[i][j-1] == c {
			return true
		}
	}
	if i < len(mat)-1 {
		if mat[i+1][j] == c {
			return true
		}
	}
	if j < len(mat[0])-1 {
		if mat[i][j+1] == c {
			return true
		}
	}
	return false
}

func Test_updateMatrix_Usage1(t *testing.T) {
	fmt.Println(updateMatrix(createTwoDimArray([]int{0, 0, 0, 0, 1, 0, 1, 1, 1}, 3, 3)))
}
