package leetcode

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	rel := map[int]int{}
	for i, v := range nums {
		if u, ok := rel[target-v]; ok {
			return []int{i, u}
		}
		rel[v] = i
	}
	return []int{}
}

func Test_twoSum_Usage(t *testing.T) {
	ret := twoSum2([]int{3, 2, 4}, 6)
	fmt.Println(ret)
}
