package leetcode

import (
	"math"
	"testing"
)

func myAtoi(s string) int {
	ind := 0
	state := 0
	flag := 1
	start_ind := 0
	end_ind := 0
	for {
		if ind == len(s) {
			if state == 1 {
				end_ind = ind
			}
			break
		}
		char := s[ind]
		if state == 0 {
			switch {
			case char == ' ':
			case char == '+':
				flag = 1
				state = 1
				start_ind = ind + 1
			case char == '-':
				flag = -1
				state = 1
				start_ind = ind + 1
			case char <= '9' && char >= '0':
				start_ind = ind
				state = 1
			default:
				state = 2
			}
			ind++
		} else if state == 1 {
			switch {
			case char <= '9' && char >= '0':
			default:
				end_ind = ind
				state = 2
			}
			ind++
		} else {
			break
		}
	}
	if end_ind < start_ind {
		end_ind = len(s)
	}
	number := 0
	MAX_N_INT := int(-math.Pow(2, 31))
	for i := start_ind; i < end_ind; i++ {
		digit := -int(s[i] - '0')
		if number < MAX_N_INT/10 {
			number = MAX_N_INT
			break
		}
		number = number * 10
		if number < MAX_N_INT-digit {
			number = MAX_N_INT
			break
		}
		number = number + digit
	}
	if flag == 1 && number == MAX_N_INT {
		return -(MAX_N_INT + 1)
	} else {
		return number * -flag
	}
}

func Test_myAtoi_Usage1(t *testing.T) {
	if myAtoi("-42") != -42 {
		t.Fail()
	}
}

func Test_myAtoi_Usage2(t *testing.T) {
	if myAtoi("-fds42") != 0 {
		t.Fail()
	}
}

func Test_myAtoi_Usage3(t *testing.T) {
	if myAtoi("42") != 42 {
		t.Fail()
	}
}

func Test_myAtoi_Usage4(t *testing.T) {
	if myAtoi("") != 0 {
		t.Fail()
	}
}

func Test_myAtoi_Usage5(t *testing.T) {
	if myAtoi("1") != 1 {
		t.Fail()
	}
}

func Test_myAtoi_Usage6(t *testing.T) {
	if myAtoi("21474836460") != 2147483647 {
		t.Fail()
	}
}
