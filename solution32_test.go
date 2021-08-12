package leetcode

import (
	"fmt"
	"testing"
)

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	start := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			start = i
			break
		}
	}
	if start == -1 {
		return 0
	}
	end := len(s)
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == ')' {
			end = i
			break
		}
	}
	if end == len(s) || end <= start {
		return 0
	}
	tmp := longestValidParentheses_valid(s[start : end+1])
	if tmp == -1 {
		return end - start + 1
	} else {
		m1 := longestValidParentheses(s[start+tmp+1 : end+1])
		m2 := longestValidParentheses(s[start : start+tmp])
		if m1 < m2 {
			return m2
		} else {
			return m1
		}
	}
}

func longestValidParentheses_valid(s string) int {
	status := 0
	status1 := 0
	for i := range s {
		if s[i] == '(' {
			status++
		} else {
			status--
		}
		if status < 0 {
			return i
		}
		if s[len(s)-1-i] == ')' {
			status1++
		} else {
			status1--
		}
		if status1 < 0 {
			return len(s) - 1 - i
		}
	}
	return -1
}

func Test_longestValidParentheses_Usage1(t *testing.T) {
	v := longestValidParentheses(")(()(()(((())(((((()()))((((()()(()()())())())()))()()()())(())()()(((()))))()((()))(((())()((()()())((())))(())))())((()())()()((()((())))))((()(((((()((()))(()()(())))((()))()))())")
	fmt.Println(v)
}

func Test_longestValidParentheses_Usage2(t *testing.T) {
	v := longestValidParentheses(")()())")
	fmt.Println(v)
}

func Test_longestValidParentheses_Usage3(t *testing.T) {
	v := longestValidParentheses("(()")
	fmt.Println(v)
}
