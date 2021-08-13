package leetcode

import (
	"testing"
)

func firstMissingPositive(nums []int) int {
	start := 0
	end := len(nums) - 1
	for {
		if start == len(nums) {
			return len(nums) + 1
		}
		if nums[start] != start+1 {
			break
		}
		start++
	}
	for {
		if end == -1 {
			return len(nums) + 1
		}
		if nums[end] != end+1 {
			break
		}
		end--
	}
	for {
		if start == end {
			break
		}
		if nums[start] == start+1 {
			start++
			continue
		}
		if nums[start] >= end+1 || nums[start] <= 0 {
			firstMissingPositive_swap(nums, start, end)
			end--
		} else {
			if nums[nums[start]-1] == nums[start] {
				nums[nums[start]-1] = 0
			}
			firstMissingPositive_swap(nums, start, nums[start]-1)
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func firstMissingPositive_swap(nums []int, a int, b int) {
	tmp := nums[a]
	nums[a] = nums[b]
	nums[b] = tmp
}

func Test_firstMissingPositive_Usage1(t *testing.T) {
	v := firstMissingPositive([]int{0, 1, 3})
	if v != 2 {
		t.FailNow()
	}
}

func Test_firstMissingPositive_Usage2(t *testing.T) {
	v := firstMissingPositive([]int{2, 2, 2})
	if v != 1 {
		t.FailNow()
	}
}
