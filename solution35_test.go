package leetcode

func searchInsert(nums []int, target int) int {
	nums = append(nums, 10000+1)
	return searchInsert_(nums, target)
}

func searchInsert_(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		}
		return -1
	}
	l := len(nums)
	r := searchInsert_(nums[:l/2], target)
	if r != -1 {
		return r
	}
	r = searchInsert_(nums[l/2:], target)
	if r != -1 {
		return r + l/2
	}
	return -1
}
