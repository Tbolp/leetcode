package leetcode

import (
	"fmt"
	"testing"
)

func beautifulArray(n int) []int {
	if n == 1 {
		return []int{1}
	}
	right := beautifulArray(n / 2)
	left := beautifulArray((n + 1) / 2)
	res := make([]int, len(left)+len(right))
	for i, v := range left {
		res[i] = 2*v - 1
	}
	for i, v := range right {
		res[i+len(left)] = 2 * v
	}
	return res
}

func Test_beautifulArray_Usage1(t *testing.T) {
	fmt.Println(beautifulArray(5))
}
