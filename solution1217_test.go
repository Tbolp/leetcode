package leetcode

func minCostToMoveChips(position []int) int {
	p1 := 0
	p2 := 0
	for _, v := range position {
		if v%2 == 0 {
			p2++
		} else {
			p1++
		}
	}
	if p1 < p2 {
		return p1
	} else {
		return p2
	}
}
