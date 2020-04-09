package leetcode

import "strconv"

func fact(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func getPermutation(n int, k int) string {
	status := make([]int, n)
	res := ""
	k = k - 1
	for i := n - 1; i >= 0; i-- {
		pos := k / fact(i)
		k = k - k/fact(i)*fact(i)
		j := -1
		for i := 0; i < n; i++ {
			if status[i] == 0 {
				j++
			}
			if j == pos {
				status[i] = 1
				res += strconv.Itoa(i + 1)
				break
			}
		}
	}
	return res
}
