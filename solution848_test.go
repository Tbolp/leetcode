package leetcode

import "testing"

func shiftingLetters(S string, shifts []int) string {
	for i := len(shifts) - 1; i >= 0; i-- {
		if i == len(shifts)-1 {
			shifts[i] = shifts[i] % 26
		} else {
			shifts[i] = (shifts[i] + shifts[i+1]) % 26
		}
	}
	s := []byte(S)
	for i, v := range s {
		s[i] = byte((((int(v) + shifts[i]) - 'a') % 26) + 'a')
	}
	return string(s)
}

func Test_shiftingLetters_Usage1(t *testing.T) {
	shiftingLetters("abc", []int{3, 5, 9})
}
