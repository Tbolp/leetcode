package leetcode

import "testing"

func longestPalindrome(s string) string {
	max_pos := 0
	max_len := 0
	state := make([]int, 2*len(s)+1)
	state[0] = 0
	for i := 0; i < len(state); i++ {
		for j := 0; j < i; j++ {
			if j+state[j]+1 == i {
				k := j - (i - j)
				if k >= 0 {
					if k%2 == 0 || s[k/2] == s[i/2] {
						state[j]++
						if state[j] > max_len {
							max_len = state[j]
							max_pos = j
						}
					}
				}
			}
		}
	}
	if max_pos%2 == 0 {
		k := max_len / 2
		i := max_pos / 2
		return s[i-k : i+k]
	} else {
		k := (max_len - 1) / 2
		i := (max_pos - 1) / 2
		return s[i-k : i+k+1]
	}
}

func Test_longestPalindrome_Usage1(t *testing.T) {
	if longestPalindrome("a") != "a" {
		t.Fail()
	}
}

func Test_longestPalindrome_Usage2(t *testing.T) {
	if longestPalindrome("babad") != "bab" {
		t.Fail()
	}
}

func Test_longestPalindrome_Usage3(t *testing.T) {
	if longestPalindrome("cbbd") != "bb" {
		t.Fail()
	}
}
