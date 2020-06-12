package leetcode

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	first := 2*numRows - 2
	second := 0
	res := ""
	for i := 0; i < numRows; i++ {
		j := i
		if j < len(s) {
			res += s[j : j+1]
		} else {
			break
		}
		for true {
			if first != 0 {
				j = j + first
				if j < len(s) {
					res += s[j : j+1]
				} else {
					break
				}
			}
			if second != 0 {
				j = j + second
				if j < len(s) {
					res += s[j : j+1]
				} else {
					break
				}
			}
		}
		first -= 2
		second += 2
	}
	return res
}
