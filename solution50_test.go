package leetcode

func myPow(x float64, n int) float64 {
	if n < 0 {
		return myPow_(1/x, n)
	}
	return myPow_(x, n)
}

func myPow_(x float64, n int) float64 {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n%2 == 0 {
		res := myPow_(x, n/2)
		return res * res
	} else {
		res := myPow_(x, n/2)
		return res * res * x
	}
}
