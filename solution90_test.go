package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	return subsetsWithDup_(nums)
}

func subsetsWithDup_(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	res := subsetsWithDup_(nums[1:])
	tmp := [][]int{}
	for _, v := range res {
		tv := make([]int, len(v))
		copy(tv, v)
		item := append(tv, nums[0])
		tmp = append(tmp, item)
	}
	for _, v := range tmp {
		flag := true
		for _, it := range res {
			is_same := true
			if len(v) == len(it) {
				for i := 0; i < len(v); i++ {
					if v[i] != it[i] {
						is_same = false
						break
					}
				}
			} else {
				is_same = false
			}
			if is_same {
				flag = false
			}
		}
		if flag {
			res = append(res, v)
		}
	}
	return res
}
