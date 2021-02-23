package leetcode

import (
	"testing"
)

func isMatch_1(s string, p string) bool {
	reg := RegularExpression{}
	res := reg.Match(s, p)
	// for _, v := range reg.cache {
	// 	fmt.Println(v)
	// }
	return res
}

type RegularExpression struct {
	cache [][]int // false -1, true 1
	multi []bool
	s     string
	p     string
}

func (this *RegularExpression) Match(s string, p string) bool {
	s = s + "0"
	tmp := ""
	this.multi = make([]bool, len(p)+1)
	for _, v := range p {
		if v != '*' {
			tmp += string(v)
		} else {
			this.multi[len(tmp)-1] = true
		}
	}
	p = tmp + "0"
	this.cache = make([][]int, len(s))
	for i := range this.cache {
		this.cache[i] = make([]int, len(p))
	}
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if p[0] == '.' {
				this.cache[0][0] = 1
			} else if p[0] == s[0] {
				this.cache[0][0] = 1
			} else {
				this.cache[0][0] = -1
			}
		} else {
			if this.multi[0] {
				if p[0] == '.' {
					this.cache[i][0] = this.cache[i-1][0]
				} else if p[0] == s[i] {
					this.cache[i][0] = this.cache[i-1][0]
				} else {
					this.cache[i][0] = -1
				}
			} else {
				this.cache[i][0] = -1
			}
		}
	}

	for i := 0; i < len(p); i++ {
		if this.multi[i] {
			if i > 0 && this.cache[0][i-1] == 1 {
				this.cache[0][i] = this.cache[0][i-1]
				continue
			}
		}
		if p[i] == '.' || p[i] == s[0] {
			flag := true
			for j := 0; j < i; j++ {
				if this.multi[j] == false {
					flag = false
				}
			}
			if !flag {
				this.cache[0][i] = -1
			} else {
				this.cache[0][i] = 1
			}
		} else {
			this.cache[0][i] = -1
		}
	}

	this.s = s
	this.p = p

	return this.getValue(len(s)-1, len(p)-1)
}

func (this *RegularExpression) getValue(x int, y int) bool {
	if this.cache[x][y] != 0 {
		return this.cache[x][y] == 1
	}
	val := false
	if this.multi[y] {
		if this.p[y] == '.' || this.p[y] == this.s[x] {
			val = this.getValue(x-1, y-1) || this.getValue(x-1, y) || this.getValue(x, y-1)
		} else {
			val = this.getValue(x, y-1)
		}
	} else {
		if this.p[y] == '.' || this.p[y] == this.s[x] {
			val = this.getValue(x-1, y-1)
		} else {
			val = false
		}
	}
	if val {
		this.cache[x][y] = 1
	} else {
		this.cache[x][y] = -1
	}
	return val
}

func Test_isMatch_1_Usage1(t *testing.T) {
	if isMatch_1("aa", "a") == true {
		t.Fail()
	}
}

func Test_isMatch_1_Usage2(t *testing.T) {
	if isMatch_1("aa", "a*") == false {
		t.Fail()
	}
}

func Test_isMatch_1_Usage3(t *testing.T) {
	if isMatch_1("ab", ".*") == false {
		t.Fail()
	}
}

func Test_isMatch_1_Usage4(t *testing.T) {
	if isMatch_1("aab", "c*a*b") == false {
		t.Fail()
	}
}

func Test_isMatch_1_Usage5(t *testing.T) {
	if isMatch_1("mississippi", "mis*is*p*.") == true {
		t.Fail()
	}
}

func Test_isMatch_1_Usage6(t *testing.T) {
	if isMatch_1("ab", ".*c") == true {
		t.Fail()
	}
}

func Test_isMatch_1_Usage7(t *testing.T) {
	if isMatch_1("aaa", "aaaa") == true {
		t.Fail()
	}
}

func Test_isMatch_1_Usage8(t *testing.T) {
	if isMatch_1("aasdfasdfasdfasdfas", "aasdf.*asdf.*asdf.*asdf.*s") == false {
		t.Fail()
	}
}
