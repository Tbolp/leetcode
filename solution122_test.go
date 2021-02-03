package leetcode

func maxProfit_1(prices []int) int {
	status := false
	buy := 0
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if status {
			if prices[i+1] < prices[i] {
				status = false
				profit += prices[i] - buy
			}
		} else {
			if prices[i+1] > prices[i] {
				status = true
				buy = prices[i]
			}
		}
	}
	if status == true {
		profit += prices[len(prices)-1] - buy
	}
	return profit
}
