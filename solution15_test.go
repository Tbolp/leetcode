package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	aSet := nums
	bSet := nums
	cSet := nums
	a := 0
	b := 0
	c := 0
	ret := [][]int{}
	for len(aSet) != 0 {
		a = aSet[0]
		bSet = aSet[1:]
		for len(bSet) != 0 {
			b = bSet[0]
			cSet = bSet[1:]
			for len(cSet) != 0 {
				c = cSet[0]
				if a+b+c == 0 {
					if len(ret) > 0 {
						if a != ret[len(ret)-1][0] ||
							b != ret[len(ret)-1][1] ||
							c != ret[len(ret)-1][2] {
							ret = append(ret, []int{a, b, c})
						}
					} else {
						ret = append(ret, []int{a, b, c})
					}
				}
				cSet = cSet[1:]
			}
			bSet = bSet[1:]
		}
		aSet = aSet[1:]
	}
	return ret
}

func Test_threeSum_Usage1(t *testing.T) {
	ret := threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6})
	fmt.Println(ret)
}

func Test_threeSum_Usage2(t *testing.T) {
	ret := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(ret)
}
