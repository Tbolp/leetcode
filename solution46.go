package leetcode

func printorder(nums []int, check []int, res *[][]int, n int) {
	if n == 0 {
		*res = append(*res, make([]int, len(nums)))
		for i, v := range check {
			(*res)[len(*res)-1][v-1] = nums[i]
		}
		return
	}
	for i := range check {
		if check[i] == 0 {
			temp := make([]int, len(check))
			copy(temp, check)
			temp[i] = n
			printorder(nums, temp, res, n-1)
		}
	}
}

func permute(nums []int) [][]int {
	res := [][]int{}
	printorder(nums, make([]int, len(nums)), &res, len(nums))
	return res
}
