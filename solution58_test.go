package leetcode

func lengthOfLastWord(s string) int {
	length := 0
	status := 0
	v := byte(0)
	for i := 0; i <= len(s); i++ {
		if i < len(s) {
			v = s[i]
		} else {
			v = ' '
		}
		switch status {
		case 0:
			if v != ' ' {
				status = 1
				length = 1
			}
		case 1:
			if v == ' ' {
				status = 0
			} else {
				length++
			}
		}
	}
	return length
}
