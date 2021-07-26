package leetcode

func majorityElement(nums []int) int {
	majorityElement_sort(nums)
	return nums[len(nums)/2+1]
}

func majorityElement_sort(nums []int) {
	if len(nums) == 0 {
		return
	}
	pivot := nums[0]
	start := 0
	end := len(nums) - 1
	is_add := false
	for start != end {
		if is_add {
			if nums[start] > pivot {
				nums[end] = nums[start]
				end--
				is_add = false
			} else {
				start++
			}
		} else {
			if nums[end] < pivot {
				nums[start] = nums[end]
				start++
				is_add = true
			} else {
				end--
			}
		}
	}
	nums[start] = pivot
	majorityElement_sort(nums[0:start])
	majorityElement_sort(nums[start+1:])
}

func majorityElement_boyer_moore(nums []int) int {
	candidate := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
