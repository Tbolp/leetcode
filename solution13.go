package leetcode

func romanToInt(s string) int {
	res := 0
	var pre byte = s[len(s)-1]
	for i := len(s) - 1; i >= 0; i-- {
		temp := 0
		switch s[i] {
		case 'I':
			if pre != 'I' {
				temp = -1
			} else {
				temp = 1
			}
		case 'V':
			if pre != 'I' && pre != 'V' {
				temp = -5
			} else {
				temp = 5
			}
		case 'X':
			if pre != 'I' && pre != 'V' && pre != 'X' {
				temp = -10
			} else {
				temp = 10
			}
		case 'L':
			if pre != 'I' && pre != 'V' && pre != 'X' && pre != 'L' {
				temp = -50
			} else {
				temp = 50
			}
		case 'C':
			if pre != 'I' && pre != 'V' && pre != 'X' && pre != 'L' && pre != 'C' {
				temp = -100
			} else {
				temp = 100
			}
		case 'D':
			if pre != 'I' && pre != 'V' && pre != 'X' && pre != 'L' && pre != 'C' && pre != 'D' {
				temp = -500
			} else {
				temp = 500
			}
		case 'M':
			if pre != 'I' && pre != 'V' && pre != 'X' && pre != 'L' && pre != 'C' && pre != 'D' && pre != 'M' {
				temp = -1000
			} else {
				temp = 1000
			}
		}
		pre = s[i]
		res += temp
	}
	return res
}
