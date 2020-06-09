package leetcode

func addBinary(a string, b string) string {
	n := len(a) - 1
	m := len(b) - 1
	res := ""
	flag := 0
	for n >= 0 || m >= 0 {
		l1 := 0
		if n >= 0 {
			l1 = int(a[n] - '0')
		}
		l2 := 0
		if m >= 0 {
			l2 = int(b[m] - '0')
		}
		l3 := l1 + l2 + flag
		switch l3 {
		case 0:
			res = "0" + res
			flag = 0
		case 1:
			res = "1" + res
			flag = 0
		case 2:
			res = "0" + res
			flag = 1
		case 3:
			res = "1" + res
			flag = 1
		}
		n--
		m--
	}
	if flag == 1 {
		res = "1" + res
	}
	return res
}
