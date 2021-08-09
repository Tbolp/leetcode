package leetcode

var minCut_Context struct {
	DP  map[string]int
	DP2 map[string]bool
}

func minCut(s string) int {
	minCut_Context.DP = map[string]int{}
	minCut_Context.DP2 = map[string]bool{}
	return minCut_(s)
}

func minCut_(s string) int {
	if s == "" {
		minCut_Context.DP[s] = -1
		return -1
	}
	if v, ok := minCut_Context.DP[s]; ok {
		return v
	}
	min := 10000
	for i := 0; i < len(s); i++ {
		if minCut_check(s[i:]) {
			tmp := minCut_(s[:i]) + 1
			if tmp < min {
				min = tmp
			}
			if min == 0 {
				return 0
			}
		}
	}
	minCut_Context.DP[s] = min
	return min
}

func minCut_check(s string) bool {
	if v, ok := minCut_Context.DP2[s]; ok {
		return v
	}
	for i := 0; i < len(s); i++ {
		j := len(s) - 1 - i
		if j <= i {
			break
		}
		if s[i] != s[j] {
			minCut_Context.DP2[s] = false
			return false
		}
	}
	minCut_Context.DP2[s] = true
	return true
}
