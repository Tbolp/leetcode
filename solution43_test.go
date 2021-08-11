package leetcode

func multiply(num1 string, num2 string) string {
	res := make([]byte, len(num1)+len(num2))
	for i := range res {
		res[i] = '0'
	}
	for i := range num1 {
		for j := range num2 {
			tmp := int((num1[i] - '0') * (num2[j] - '0'))
			ind := i + j + 1
			for tmp != 0 {
				tmp = int(res[ind]-'0') + tmp
				res[ind] = byte(tmp%10) + '0'
				tmp = tmp / 10
				ind--
			}
		}
	}
	for res[0] == '0' && len(res) > 1 {
		res = res[1:]
	}
	return string(res)
}
