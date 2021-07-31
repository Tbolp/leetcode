package leetcode

func firstUniqChar(s string) int {
	alpha := make([]int, 26)
	for i := range s {
		alpha[s[i]-'a']++
	}
	for i := range s {
		if alpha[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
