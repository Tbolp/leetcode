package leetcode

import (
	"testing"
)

func minFlipsMonoIncr(s string) int {
	if len(s) < 2 {
		return 0
	}
	ones := make([]int, len(s))
	dp := make([]int, len(s))
	ones[0] = int(s[0] - '0')
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if ones[i-1] < 1+dp[i-1] {
				dp[i] = ones[i-1]
			} else {
				dp[i] = 1 + dp[i-1]
			}
			ones[i] = ones[i-1]
		} else {
			if 1+ones[i-1] < dp[i-1] {
				dp[i] = ones[i-1] + 1
			} else {
				dp[i] = dp[i-1]
			}
			ones[i] = ones[i-1] + 1
		}
	}
	return dp[len(dp)-1]
}

func Test_minFlipsMonoIncr_Usage1(t *testing.T) {
	minFlipsMonoIncr("00110")
}
