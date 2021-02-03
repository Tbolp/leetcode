package leetcode

import (
	"testing"
)

func maxSubArray_old(nums []int) int {
	l := 0
	r := 1
	max := sumSubArray(nums[l:r])
	end := len(nums)
	for r := 1; r <= end; r++ {
		temp := sumSubArray(nums[l:r])
		if temp > max {
			max = temp
		}
		for j := l; j < r-1; j++ {
			temp = temp - nums[j]
			if temp > max {
				max = temp
				l = j
			}
		}
	}
	return max
}

func sumSubArray(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	max := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		temp := dp[i-1] + nums[i]
		if temp < nums[i] {
			dp[i] = nums[i]
		} else {
			dp[i] = temp
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func Test_maxSubArray_Usage1(t *testing.T) {
	if maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}) != 6 {
		t.Fail()
	}
}

func Test_maxSubArray_Usage2(t *testing.T) {
	if maxSubArray([]int{8, -19, 5, -4, 20}) != 21 {
		t.Fail()
	}
}

func Test_maxSubArray_Usage3(t *testing.T) {
	if maxSubArray([]int{-1}) != -1 {
		t.Fail()
	}
}

func Test_maxSubArray_Usage4(t *testing.T) {
	if maxSubArray([]int{1, 2}) != 3 {
		t.Fail()
	}
}
