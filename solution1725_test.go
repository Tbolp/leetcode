package leetcode

func countGoodRectangles(rectangles [][]int) int {
	count := 0
	maxLen := 0
	for _, v := range rectangles {
		l := minInt(v[0], v[1])
		if l > maxLen {
			count = 1
			maxLen = l
		} else if l == maxLen {
			count++
		}
	}
	return count
}
