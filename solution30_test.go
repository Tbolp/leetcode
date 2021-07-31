package leetcode

func findSubstring(s string, words []string) []int {
	wordlen := len(words[0])
	res := []int{}
	for i := 0; i < len(s); i++ {
		if len(s)-i < len(words)*wordlen {
			return res
		}
		if findSubstring_check(s[i:], words) {
			res = append(res, i)
		}
	}
	return res
}

func findSubstring_check(s string, words []string) bool {
	wordlen := len(words[0])
	tmp := map[string]int{}
	for _, v := range words {
		tmp[v]++
	}
	for i := 0; i+wordlen <= len(s); i += wordlen {
		v := s[i : i+wordlen]
		if _, ok := tmp[v]; ok {
			tmp[v]--
			if tmp[v] == 0 {
				delete(tmp, v)
			}
		} else {
			return false
		}
		if len(tmp) == 0 {
			return true
		}
	}
	return false
}
