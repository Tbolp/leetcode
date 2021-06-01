package leetcode

import (
	"math"
	"sort"
	"testing"
)

func threeSumClosest(nums []int, target int) int {
	index := [3]int{}
	sort.Ints(nums)
	res := math.MaxInt32
	detal := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		index[0] = i
		index[1] = i + 1
		index[2] = len(nums) - 1
		tmp := threeSumClosest_(nums, target, index)
		tmp_detal := tmp - target
		if tmp_detal < 0 {
			tmp_detal = -tmp_detal
		}
		if tmp_detal == 0 {
			return target
		}
		if tmp_detal < detal {
			detal = tmp_detal
			res = tmp
		}
	}
	return res
}

func threeSumClosest_(nums []int, target int, index [3]int) int {
	res := math.MaxInt32
	detal := math.MaxInt32
	for index[1] < index[2] {
		tmp := nums[index[0]] + nums[index[1]] + nums[index[2]]
		if tmp == target {
			return target
		} else if tmp < target {
			index[1]++
		} else {
			index[2]--
		}
		tmp_detal := tmp - target
		if tmp_detal < 0 {
			tmp_detal = -tmp_detal
		}
		if tmp_detal < detal {
			detal = tmp_detal
			res = tmp
		}
	}
	return res
}

func Test_threeSumClosest_Usage(t *testing.T) {
	if threeSumClosest([]int{-1, 2, 1, -4}, 1) != 2 {
		t.Fail()
	}
}

func Test_threeSumClosest_Usage2(t *testing.T) {
	if threeSumClosest([]int{-1, 2, 1, 0, 0, 0}, 1) != 1 {
		t.Fail()
	}
}

func Test_threeSumClosest_Usage3(t *testing.T) {
	if threeSumClosest([]int{1, 6, 9, 14, 16, 70}, 81) != 80 {
		t.Fail()
	}
}
