package leetcode

import (
	"math"
	"testing"
)

func longestConsecutive(nums []int) int {
	tree := make([]int, len(nums))
	rel := map[int]int{}
	for i := range tree {
		if _, ok := rel[nums[i]]; ok {
			tree[i] = rel[nums[i]]
		} else {
			tree[i] = -1
			rel[nums[i]] = i
		}
	}
	max_len := 0
	if len(nums) > 0 {
		max_len = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if math.Abs(float64(nums[i]-nums[j])) <= 1 {
				k := i
				for tree[k] >= 0 {
					k = tree[k]
				}
				l := j
				for tree[l] >= 0 {
					l = tree[l]
				}
				if k == l {
					continue
				}
				tree[k] += tree[l]
				if -tree[k] > max_len {
					max_len = -tree[k]
				}
				tree[l] = k
			}
		}
	}
	return max_len
}

func Test_longestConsecutive_Usage1(t *testing.T) {
	longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
}
