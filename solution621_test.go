package leetcode

import (
	"testing"
)

func leastInterval(tasks []byte, n int) int {
	rel := map[byte]int{}
	maxop := 0
	for _, v := range tasks {
		rel[v]++
		if rel[v] > maxop {
			maxop = rel[v]
		}
	}
	res := (n + 1) * (maxop - 1)
	for _, v := range rel {
		if v == maxop {
			res++
		}
	}
	if res < len(tasks) {
		res = len(tasks)
	}
	return res
}

func Test_leastInterval_Usage1(t *testing.T) {
	if leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B', 'B'}, 2) != 10 {
		t.Fail()
	}
}

func Test_leastInterval_Usage2(t *testing.T) {
	if leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}, 2) != 12 {
		t.Fail()
	}
}
