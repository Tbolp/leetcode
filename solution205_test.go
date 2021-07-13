package leetcode

import "testing"

func isIsomorphic(s string, t string) bool {
	if isIsomorphic_(s, t) && isIsomorphic_(t, s) {
		return true
	}
	return false
}

func isIsomorphic_(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	fmap := map[byte]byte{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if _, ok := fmap[ch]; !ok {
			fmap[ch] = t[i]
		} else {
			if fmap[ch] != t[i] {
				return false
			}
		}
	}
	return true
}

func Test_isIsomorphic_Usage1(t *testing.T) {
	if isIsomorphic("badc", "baba") {
		t.Fail()
	}
}
