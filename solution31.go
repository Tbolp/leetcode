package leetcode

import "sort"

func swapnearest(nums []int) bool {
	k := 0
	flag := false
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[0] {
			flag = true
			if k == 0 {
				k = i
			} else if nums[i]-nums[0] < nums[k]-nums[0] {
				k = i
				flag = true
			}
		}
	}
	if flag {
		temp := nums[0]
		nums[0] = nums[k]
		nums[k] = temp
	}
	return flag
}

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if swapnearest(nums[i:]) {
			sort.Ints(nums[i+1:])
			return
		}
	}
	sort.Ints(nums)
}
