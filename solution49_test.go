package leetcode

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	for i := range strs {
		if strs[i] == "0" {
			continue
		}
		tmp := []string{}
		tmp = append(tmp, strs[i])
		for j := i + 1; j < len(strs); j++ {
			if strs[j] == "0" {
				continue
			}
			if groupAnagrams_equal(strs[i], strs[j]) {
				tmp = append(tmp, strs[j])
				strs[j] = "0"
			}
		}
		res = append(res, tmp)
	}
	return res
}

func groupAnagrams_equal(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	slot := make([]int, 26)
	for i := range str1 {
		slot[str1[i]-'a']++
		slot[str2[i]-'a']--
	}
	for i := range slot {
		if slot[i] != 0 {
			return false
		}
	}
	return true
}
