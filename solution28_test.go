package leetcode

import "testing"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := make([]int, len(needle))
	next[0] = 0
	for i := 1; i < len(needle); i++ {
		next[i] = 0
		j := i - 1
		for j > -1 {
			if needle[i] == needle[next[j]] {
				next[i] = next[j] + 1
				break
			} else {
				j = next[j] - 1
			}
		}
	}
	for i := len(next) - 1; i > 0; i-- {
		next[i] = next[i-1]
	}
	next[0] = -1
	for i := 1; i < len(next); i++ {
		if needle[i] == needle[next[i]] {
			next[i] = next[next[i]]
		}
	}
	j := 0
	for i := 0; i < len(haystack); {
		if haystack[i] == needle[j] {
			i++
			j++
		} else if next[j] == -1 {
			i++
			j = 0
		} else {
			j = next[j]
		}
		if j == len(needle) {
			return i - len(needle)
		}
	}
	return -1
}

func Test_strStr_Usage1(t *testing.T) {
	if strStr("mississippi", "issip") != 4 {
		t.FailNow()
	}
}

func Test_strStr_Usage2(t *testing.T) {
	if strStr("babba", "bbb") != -1 {
		t.FailNow()
	}
}
