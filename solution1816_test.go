package leetcode

func truncateSentence(s string, k int) string {
	state := 0
	for i := range s {
		if s[i] != ' ' {
			if state <= 0 {
				state = -state
				state++
			}
		} else {
			if state == k {
				return s[0:i]
			}
			state = -state
		}
	}
	return s
}
