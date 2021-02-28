package leetcode

func letterCombinations(digits string) []string {
	rel := map[byte]string{}
	rel['2'] = "abc"
	rel['3'] = "def"
	rel['4'] = "ghi"
	rel['5'] = "jkl"
	rel['6'] = "mno"
	rel['7'] = "pqrs"
	rel['8'] = "tuv"
	rel['9'] = "wxyz"
	res := []string{}
	if digits == "" {
		return res
	}
	letterCombinations_(digits, rel, "", &res)
	return res
}
func letterCombinations_(current string, rel map[byte]string, str string, res *[]string) {
	if len(current) == 0 {
		*res = append(*res, str)
		return
	}
	letter := current[0]
	for i := 0; i < len(rel[letter]); i++ {
		letterCombinations_(current[1:], rel, str+string(rel[letter][i]), res)
	}
}
