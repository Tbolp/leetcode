package leetcode

import (
	"testing"
)

func divisorGame(N int) bool {
	dp := make([]bool, N+1)
	dp[1] = false
	for i := 2; i <= N; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 {
				if dp[i-j] == false {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[N]
}

func Test_divisorGame_Usage1(t *testing.T) {
	if divisorGame(1) == true {
		t.Fail()
	}
}
