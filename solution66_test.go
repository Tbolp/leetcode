package leetcode

func plusOne(digits []int) []int {
	cur := len(digits) - 1
	flag := true
	for flag && cur >= 0 {
		flag = false
		digits[cur] += 1
		if digits[cur]/10 != 0 {
			flag = true
		} else {
			break
		}
		digits[cur] = digits[cur] % 10
		cur--
	}
	if flag {
		digits = append([]int{1}, digits...)
	}
	return digits
}
