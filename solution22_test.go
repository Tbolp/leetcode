package leetcode

import (
	"fmt"
	"testing"
)

func generateParenthesis(n int) []string {
	res := []string{}
	generateParenthesis_(n, 0, "", &res)
	return res
}

func generateParenthesis_(remain int, current int, str string, res *[]string) {
	if remain == 0 && current == 0 {
		*res = append(*res, str)
		return
	}
	if remain == 0 {
		generateParenthesis_(remain, current-1, str+")", res)
	} else {
		generateParenthesis_(remain-1, current+1, str+"(", res)
		if current > 0 {
			generateParenthesis_(remain, current-1, str+")", res)
		}
	}
}

func Test_generateParenthesis_Usage1(t *testing.T) {
	res := generateParenthesis(3)
	fmt.Println(res)
}
