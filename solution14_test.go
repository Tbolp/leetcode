package leetcode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ind := 0
	for {
		flag := true
		for _, str := range strs {
			if ind >= len(str) {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		chr := strs[0][ind]
		for _, str := range strs {
			if str[ind] != chr {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		ind++
	}
	return strs[0][0:ind]
}
