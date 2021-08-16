package leetcode

import (
	"fmt"
	"testing"
)

func minWindow(s string, t string) string {
	cache := map[byte]int{}
	min := map[byte]int{}
	s = t + s
	dp := make([]int, len(s))
	dp[0] = len(t)
	i := 0
	j := len(t)
	for i := range t {
		cache[t[i]]++
		min[t[i]]++
	}
	for {
		if _, ok := cache[s[i]]; ok {
			cache[s[i]]--
			if cache[s[i]] >= min[s[i]] {
				dp[i+1] = dp[i] - 1
			} else {
				for j < len(s) && s[j] != s[i] {
					if _, ok := cache[s[j]]; ok {
						cache[s[j]]++
					}
					j++
				}
				if j == len(s) {
					break
				} else {
					dp[i+1] = j - i
					cache[s[j]]++
					j++
				}
			}
		} else {
			dp[i+1] = dp[i] - 1
		}
		i++
		if i == len(dp)-1 {
			break
		}
	}
	pos := len(t)
	for i := len(t); i < len(s); i++ {
		if dp[i] != 0 && dp[i] < dp[pos] {
			pos = i
		}
	}
	return s[pos : pos+dp[pos]]
}

func Test_minWindow_Usage1(t *testing.T) {
	v := minWindow("ADOBECODEBANC", "ABC")
	fmt.Println(v)
}

func Test_minWindow_Usage2(t *testing.T) {
	v := minWindow("a", "aa")
	fmt.Println(v)
}

func Test_minWindow_Usage3(t *testing.T) {
	v := minWindow("a", "a")
	fmt.Println(v)
}

func Test_minWindow_Usage4(t *testing.T) {
	v := minWindow("acbdbaab", "aabd")
	fmt.Println(v)
}
