package leetcode

func addStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	tmp := byte(0)
	flag := byte(0)
	res := ""
	for i >= 0 || j >= 0 {
		if i >= 0 && j >= 0 {
			tmp = num1[i] - '0' + num2[j] - '0' + flag
			flag = tmp / 10
			tmp = tmp % 10
			res = string(tmp+'0') + res
		} else if i >= 0 {
			tmp = num1[i] - '0' + flag
			flag = tmp / 10
			tmp = tmp % 10
			res = string(tmp+'0') + res
		} else {
			tmp = num2[j] - '0' + flag
			flag = tmp / 10
			tmp = tmp % 10
			res = string(tmp+'0') + res
		}
		i--
		j--
	}
	if flag == 1 {
		res = "1" + res
	}
	return res
}
