package leetcode

func lengthOfLongestSubstring(s string) int {
	counts := make([]int, len(s))
	res := 0
	j := 0
	for i := range s {
		if i != 0 {
			counts[i] = counts[i-1] - 1
		}
		if j == i {
			counts[i]++
			j++
		}
		for j < len(s) {
			flag := false
			for k := i; k < j; k++ {
				if s[k] == s[j] {
					flag = true
					break
				}
			}
			if flag {
				break
			}
			counts[i]++
			j++
		}
		if counts[i] > res {
			res = counts[i]
		}
	}
	return res
}
