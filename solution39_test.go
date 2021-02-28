package leetcode

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	combinationSum_(candidates, target, []int{}, &res)
	return res
}

func combinationSum_(candidates []int, target int, set []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, set)
		return
	}
	if target < 0 {
		return
	}
	for i := 0; i < len(candidates); i++ {
		tmp := make([]int, len(set))
		copy(tmp, set)
		tmp = append(tmp, candidates[i])
		combinationSum_(candidates[i:], target-candidates[i], tmp, res)
	}
}
