package leetcode

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	for i := 1; i <= x; i++ {
		if i*i > x {
			return i - 1
		}
	}
	return x
}
