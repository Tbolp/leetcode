package leetcode

import (
	"strconv"
)

func readBinaryWatch(num int) []string {
	clock := []int{8, 4, 2, 1, 32, 16, 8, 4, 2, 1}
	res := []string{}
	readBinaryWatch_(num, -1, clock, 0, 0, &res)
	return res
}

func readBinaryWatch_(remain int, current int, clock []int, hour int, minute int, res *[]string) {
	if remain == 0 {
		if hour >= 0 && hour < 12 && minute >= 0 && minute < 60 {
			show := ""
			show += strconv.Itoa(hour) + ":"
			if minute < 10 {
				show += "0" + strconv.Itoa(minute)
			} else {
				show += strconv.Itoa(minute)
			}
			*res = append(*res, show)
		}
		return
	}
	for i := current + 1; i < len(clock); i++ {
		if i < 4 {
			readBinaryWatch_(remain-1, i, clock, hour+clock[i], minute, res)
		} else {
			readBinaryWatch_(remain-1, i, clock, hour, minute+clock[i], res)
		}
	}
}
