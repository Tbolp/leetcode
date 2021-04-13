package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	digits := []int{}
	for x != 0 {
		digits = append(digits, x%10)
		x = x / 10
	}
	for i := 0; i < len(digits); i++ {
		j := len(digits) - 1 - i
		if j < i {
			return true
		}
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}
