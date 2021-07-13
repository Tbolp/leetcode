package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func fourSum(nums []int, target int) [][]int {
	app := nSum{}
	app.res = make([][]int, 0)
	app.tmp = make([]int, 0)
	sort.Ints(nums)
	app.do(nums, target, 4)
	return app.res
}

type nSum struct {
	res [][]int
	tmp []int
}

func (this *nSum) do(nums []int, target int, count int) {
	if count == 1 {
		for _, v := range nums {
			if target == v {
				this.tmp = append(this.tmp, target)
				res := make([]int, len(this.tmp))
				copy(res, this.tmp)
				this.res = append(this.res, res)
				this.tmp = this.tmp[:len(this.tmp)-1]
				return
			}
		}
		return
	}
	for i := 0; i < len(nums) && len(nums) >= count; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if target < 0 && nums[0] > 0 {
			continue
		}
		if target > 0 && target < nums[0] {
			continue
		}
		this.tmp = append(this.tmp, nums[i])
		this.do(nums[i+1:], target-nums[i], count-1)
		this.tmp = this.tmp[:len(this.tmp)-1]
	}
}

func Test_fourSum_Usage1(t *testing.T) {
	res := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println(res)
	if len(res) != 3 {
		t.Fail()
	}
}

func Test_fourSum_Usage2(t *testing.T) {
	res := fourSum([]int{2, 2, 2, 2, 2}, 8)
	fmt.Println(res)
	if len(res) != 1 {
		t.Fail()
	}
}

func Test_fourSum_Usage3(t *testing.T) {
	res := fourSum([]int{-2, -1, -1, 1, 1, 2, 2}, 0)
	fmt.Println(res)
	if len(res) != 2 {
		t.Fail()
	}
}

func Test_fourSum_Usage4(t *testing.T) {
	res := fourSum([]int{-493, -482, -482, -456, -427, -405, -392, -385, -351, -269, -259, -251, -235, -235, -202, -201, -194, -189, -187, -186, -180, -177, -175, -156, -150, -147, -140, -122, -112, -112, -105, -98, -49, -38, -35, -34, -18, 20, 52, 53, 57, 76, 124, 126, 128, 132, 142, 147, 157, 180, 207, 227, 274, 296, 311, 334, 336, 337, 339, 349, 354, 363, 372, 378, 383, 413, 431, 471, 474, 481, 492}, 6189)
	fmt.Println(res)
	if len(res) != 0 {
		t.Fail()
	}
}

func Test_fourSum_Usage5(t *testing.T) {
	res := fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11)
	fmt.Println(res)
	if len(res) != 1 {
		t.Fail()
	}
}
