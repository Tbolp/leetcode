package leetcode

import (
	"testing"
)

func isMatch(s string, p string) bool {
	np := ""
	for i := 0; i < len(p); i++ {
		if len(np) == 0 || !(np[len(np)-1] == '*' && p[i] == '*') {
			np += string(p[i])
		}
	}
	s = s + "0"
	p = np + "0"
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(p))
	}
	status := 0
	for j := 0; j < len(p); j++ {
		switch status {
		case 0:
			if p[j] == '*' {
				dp[0][j] = 1
			} else if p[j] == '?' || p[j] == s[0] {
				dp[0][j] = 1
				status = 1
			} else {
				dp[0][j] = -1
				status = 1
			}
		case 1:
			if p[j] == '*' {
				dp[0][j] = dp[0][j-1]
			} else {
				dp[0][j] = -1
			}
		}
	}
	for i := 0; i < len(s); i++ {
		if p[0] == '*' {
			dp[i][0] = 1
		} else if i == 0 {
			if p[0] == s[i] || p[0] == '?' {
				dp[i][0] = 1
			} else {
				dp[i][0] = -1
			}
		} else {
			dp[i][0] = -1
		}
	}
	return isMatchCache(s, p, dp)
}

func isMatchCache(s string, p string, dp [][]int) bool {
	temp := dp[len(s)-1][len(p)-1]
	if temp != 0 {
		return temp == 1
	}
	c := p[len(p)-1]
	i := len(s) - 1
	j := len(p) - 1
	v := false
	switch c {
	case '*':
		v1 := isMatchCache(s[:i], p, dp)
		v2 := isMatchCache(s, p[:j], dp)
		v = v1 || v2
	default:
		if c != '?' {
			if s[i] != c {
				dp[i][j] = -1
				return false
			}
		}
		v = isMatchCache(s[:i], p[:j], dp)
	}
	if v {
		dp[i][j] = 1
	} else {
		dp[i][j] = -1
	}
	return v
}

func Test_isMatch_Usage1(t *testing.T) {
	if isMatch("aa", "a") == true {
		t.Fail()
	}
}

func Test_isMatch_Usage2(t *testing.T) {
	if isMatch("aa", "*") == false {
		t.Fail()
	}
}

func Test_isMatch_Usage3(t *testing.T) {
	if isMatch("cb", "?a") == true {
		t.Fail()
	}
}

func Test_isMatch_Usage4(t *testing.T) {
	if isMatch("adceb", "*a*b") == false {
		t.Fail()
	}
}

func Test_isMatch_Usage5(t *testing.T) {
	if isMatch("acdcb", "a*c?b") == true {
		t.Fail()
	}
}

func Test_isMatch_Usage6(t *testing.T) {
	if isMatch("", "*****") == false {
		t.Fail()
	}
}

func Test_isMatch_Usage7(t *testing.T) {
	if isMatch("bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab", "b*b*ab**ba*b**b***bba") == true {
		t.Fail()
	}
}

func Test_isMatch_Usage8(t *testing.T) {
	if isMatch("a", "aa") == true {
		t.Fail()
	}
}

func Test_isMatch_Usage9(t *testing.T) {
	if isMatch("b", "?*?") == true {
		t.Fail()
	}
}

func Test_isMatch_Usage10(t *testing.T) {
	if isMatch("abaabaaaabbabbaaabaabababbaabaabbabaaaaabababbababaabbabaabbbbaabbbbbbbabaaabbaaaaabbaabbbaaaaabbbabb", "ab*aaba**abbaaaa**b*b****aa***a*b**ba*a**ba*baaa*b*ab*") {
		t.Fail()
	}
}

func Test_isMatch_Usage11(t *testing.T) {
	if isMatch("aab", "c*a*b") {
		t.Fail()
	}
}
