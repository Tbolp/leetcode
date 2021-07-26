package leetcode

func partitionDisjoint(nums []int) int {
	ind := 0
	max := nums[ind]
	tmpmax := nums[ind]
	for i := 1; i < len(nums); i++ {
		if nums[i] >= max {
			if nums[i] > tmpmax {
				tmpmax = nums[i]
			}
			continue
		} else {
			ind = i
			max = tmpmax
		}
	}
	return ind + 1
}
