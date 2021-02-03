package leetcode

import (
	"fmt"
	"testing"
)

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < len(dp); i++ {
		dp[i] = minInt(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}
	fmt.Println(dp)
	return dp[len(cost)]
}

func Test_minCostClimbingStairs_Usage1(t *testing.T) {
	minCostClimbingStairs([]int{10, 15, 20})
}
