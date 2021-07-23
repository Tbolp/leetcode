package leetcode

func partitionDisjoint(nums []int) int {
	min_ind := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[min_ind] {
			min_ind = i
		}
	}
	if min_ind == 0 {
		return 1
	}
	max := nums[0]
	for i := 0; i < min_ind; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < max {
			return i + 1
		}
	}
	return -1
}
