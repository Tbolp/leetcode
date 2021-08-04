package leetcode

func jump(nums []int) int {
	kno := make([]int, len(nums))
	return jump_(nums, kno)
}

func jump_(nums []int, kno []int) int {
	if len(nums) == 1 {
		return 0
	}
	if kno[len(nums)-1] != 0 {
		return kno[len(nums)-1]
	}
	min := -1
	for i := 0; i < len(nums)-1; i++ {
		pre := jump_(nums[:i+1], kno)
		step := nums[i] + i
		if pre != -1 && step >= len(nums)-1 {
			if min == -1 || pre+1 < min {
				min = pre + 1
			}
		}
	}
	kno[len(nums)-1] = min
	return kno[len(nums)-1]
}

func jump1(nums []int) int {
	kno := make([]int, len(nums))
	for i := 0; i < len(kno); i++ {
		kno[i] = 100000
	}
	kno[0] = 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j]+j >= i && kno[j]+1 < kno[i] {
				kno[i] = kno[j] + 1
			}
		}
	}
	return kno[len(nums)-1]
}
