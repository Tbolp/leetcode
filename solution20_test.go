package leetcode

func isValid(s string) bool {
	stack := make([]byte, len(s))
	stack = stack[0:0]
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		} else {
			top := stack[len(stack)-1]
			switch s[i] {
			case '(':
				fallthrough
			case '[':
				fallthrough
			case '{':
				stack = append(stack, s[i])
			case ')':
				if top != '(' {
					return false
				}
				stack = stack[0 : len(stack)-1]
			case ']':
				if top != '[' {
					return false
				}
				stack = stack[0 : len(stack)-1]
			case '}':
				if top != '{' {
					return false
				}
				stack = stack[0 : len(stack)-1]
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
