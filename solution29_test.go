package leetcode

import (
	"math"
	"testing"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	flag := -1
	if dividend > 0 {
		dividend = -dividend
		flag = 1
	}
	if divisor > 0 {
		divisor = -divisor
	} else {
		flag = flag * -1
	}
	remain := dividend
	count := 0
	for remain <= divisor {
		count++
		remain = remain - divisor
	}
	count = flag * count
	if count > math.MaxInt32 {
		return math.MaxInt32
	}
	return count
}

func Test_divide_Usage1(t *testing.T) {
	if divide(-2147483648, -1) !=
		2147483647 {
		t.FailNow()
	}
}
