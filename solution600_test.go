package leetcode

func findIntegers(n int) int {
	order := 1
	total := 1
	for total >= n {
		order++
		total = (total+1)*2 - 1
	}
	return findIntegers_(order, n)
}

func findIntegers_(n int, k int) int {
	if k == 0 {
		return 1
	}
	return findIntegers_(n-2, (k-1)/4) + findIntegers_(n-1, k/2)
}
