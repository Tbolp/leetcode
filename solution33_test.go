package leetcode

import "testing"

func search(nums []int, target int) int {
	min := search_max_min(nums)
	end := min - 1
	if end < 0 {
		end = len(nums) - 1
	}
	return search_(nums, min, end, target)
}

func search_max_min(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	llen := len(nums) / 2
	lmin := search_max_min(nums[:llen])
	rmin := search_max_min(nums[llen:])
	if nums[lmin] < nums[llen+rmin] {
		return lmin
	} else {
		return llen + rmin
	}
}

func search_(nums []int, start int, end int, target int) int {
	mid := 0
	if end < start {
		mid = (start + end + len(nums)) / 2 % len(nums)
	} else {
		mid = (start + end) / 2
	}
	if target == nums[mid] {
		return mid
	} else if start == end {
		return -1
	} else if target < nums[mid] && start != mid {
		return search_(nums, start, (mid-1+len(nums))%len(nums), target)
	} else if end != mid {
		return search_(nums, (mid+1)%len(nums), end, target)
	}
	return -1
}

func Test_search_Usage1(t *testing.T) {
	search([]int{8, 9, 0, 1, 3, 4, 6}, 7)
}
