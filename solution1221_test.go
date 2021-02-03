package leetcode

func balancedStringSplit(s string) int {
	balance := 0
	count := 0
	for i := range s {
		if s[i] == 'R' {
			balance++
		} else {
			balance--
		}
		if balance == 0 {
			count++
		}
	}
	return count
}
