package leetcode

import (
	"fmt"
	"math"
	"testing"
)

func maxProfit(prices []int) int {
	l := 0
	r := 0
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] > prices[r] {
			r = i
			profit = int(math.Max(float64(prices[r]-prices[l]), float64(profit)))
		} else if prices[i] < prices[l] {
			l = i
			r = l
		}
	}
	return profit
}

func Test_maxProfit_Usage1(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func Test_maxProfit_Usage2(t *testing.T) {
	fmt.Println(maxProfit([]int{3, 2, 6, 5, 0, 3}))
}
