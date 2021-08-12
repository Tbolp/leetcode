package leetcode

import (
	"fmt"
	"testing"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	res := []int{}
	start := 0
	end := len(nums)
	for {
		if end == start {
			if nums[0] == target {
				res = append(res, 0)
				break
			} else {
				return []int{-1, -1}
			}
		}
		check := (end + start) / 2
		if nums[check] >= target {
			end = check
		} else {
			if len(nums) > check+1 && nums[check+1] == target {
				res = append(res, check+1)
				break
			} else {
				start = check + 1
			}
		}
	}
	start = 0
	end = len(nums)
	for {
		if end == start {
			if nums[len(nums)-1] == target {
				res = append(res, len(nums)-1)
				break
			} else {
				return []int{-1, -1}
			}
		}
		check := (end + start) / 2
		if nums[check] <= target {
			start = check + 1
		} else {
			if check-1 >= 0 && nums[check-1] == target {
				res = append(res, check-1)
				break
			} else {
				end = check
			}
		}
	}
	return res
}

func Test_searchRange_Usage1(t *testing.T) {
	v := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println(v)
}

func Test_searchRange_Usage2(t *testing.T) {
	v := searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	fmt.Println(v)
}

func Test_searchRange_Usage3(t *testing.T) {
	v := searchRange([]int{1}, 1)
	fmt.Println(v)
}

func Test_searchRange_Usage4(t *testing.T) {
	v := searchRange([]int{2, 2}, 1)
	fmt.Println(v)
}
