package leetcode

func sortColors(nums []int) {
	rel := map[int]int{}
	rel[0] = 0
	rel[1] = 0
	rel[2] = 0
	for _, val := range nums {
		rel[val]++
	}
	ind := 0
	for i := 0; i < rel[0]; i++ {
		nums[ind] = 0
		ind++
	}
	for i := 0; i < rel[1]; i++ {
		nums[ind] = 1
		ind++
	}
	for i := 0; i < rel[2]; i++ {
		nums[ind] = 2
		ind++
	}
}
