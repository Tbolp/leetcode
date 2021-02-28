package leetcode

import (
	"fmt"
	"testing"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	combinationSum2_(candidates, target, map[int]interface{}{}, []int{}, &res)
	return res
}

func combinationSum2_(candidates []int, target int, forbid map[int]interface{}, set []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, set)
		return
	}
	if target < 0 {
		return
	}
	for i := 0; i < len(candidates); i++ {
		if _, ok := forbid[candidates[i]]; ok {
			continue
		}
		tmp := make([]int, len(set))
		copy(tmp, set)
		tmp = append(tmp, candidates[i])
		forbidtmp := map[int]interface{}{}
		for k, v := range forbid {
			forbidtmp[k] = v
		}
		combinationSum2_(candidates[i+1:], target-candidates[i], forbidtmp, tmp, res)
		forbid[candidates[i]] = nil
	}
}

func Test_combinationSum2_Usage1(t *testing.T) {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
