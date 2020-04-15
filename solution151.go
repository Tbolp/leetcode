package leetcode

func reverseWords(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			e := i
			for i < len(s) {
				if s[i] == ' ' {
					break
				}
				i++
			}
			res = s[e:i] + " " + res
		}
	}
	if len(res) == 0 {
		return res
	}
	if res[len(res)-1] == ' ' {
		return res[0 : len(res)-1]
	}
	return res
}
