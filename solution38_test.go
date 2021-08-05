package leetcode

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	tmp := countAndSay(n - 1)
	res := ""
	cnt := 0
	cur := tmp[0]
	for i := range tmp {
		if tmp[i] == cur {
			cnt++
		} else {
			res += strconv.Itoa(cnt) + string(cur)
			cnt = 1
			cur = tmp[i]
		}
	}
	res += strconv.Itoa(cnt) + string(cur)
	return res
}
