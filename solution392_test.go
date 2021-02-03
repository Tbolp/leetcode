package leetcode

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	i := 0
	for _, v := range t {
		if rune(s[i]) == v {
			i++
			if i == len(s) {
				return true
			}
		}
	}
	return false
}
