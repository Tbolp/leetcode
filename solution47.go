package leetcode

func printorderunique(nums []int, check []int, res *[][]int, n int) {
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
			flag := true
			for j := i + 1; j < len(check); j++ {
				if nums[j] == nums[i] && check[j] != 0 {
					flag = false
					break
				}
			}
			if flag == true {
				printorderunique(nums, temp, res, n-1)
			}
		}
	}
}

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	printorderunique(nums, make([]int, len(nums)), &res, len(nums))
	return res
}
