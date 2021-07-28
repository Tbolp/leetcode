package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	end := 0
	cur := nums[end]
	count := 1
	for i, v := range nums {
		if i == 0 {
			continue
		}
		if v != cur {
			cur = v
			end++
			count++
			nums[end] = v
		}
	}
	return count
}
