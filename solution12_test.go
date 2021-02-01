package leetcode

import (
	"strings"
)

func intToRoman(num int) string {
	unit := []string{"I", "X", "C", "M"}
	center := []string{"V", "L", "D"}
	r := 0
	ret := ""
	i := 0
	for num != 0 {
		r = num % 10
		num = num / 10
		if r < 4 {
			ret = strings.Repeat(unit[i], r) + ret
		} else if r == 4 {
			ret = unit[i] + center[i] + ret
		} else if r == 5 {
			ret = center[i] + ret
		} else if r < 9 {
			ret = center[i] + strings.Repeat(unit[i], r-5) + ret
		} else if r == 9 {
			ret = unit[i] + unit[i+1] + ret
		}
		i++
	}
	return ret
}
