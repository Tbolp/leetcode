package leetcode

import (
	"math"
	"testing"
)

func reverse(x int) int {
	y := 0
	flag := 1
	if x > 0 {
		x = -x
		flag = -1
	}
	N_MAX_INT := int(-math.Pow(2, 31))
	for x != 0 {
		if y < N_MAX_INT/10 {
			return 0
		}
		y = y * 10
		if y < N_MAX_INT-x%10 {
			return 0
		}
		y = y + x%10
		x = x / 10
	}
	if flag == -1 && y == N_MAX_INT {
		return 0
	}
	return y * flag
}

func Test_reverse_Usage1(t *testing.T) {
	if reverse(123) != 321 {
		t.Fail()
	}
}
