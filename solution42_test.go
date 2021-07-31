package leetcode

func trap_brute(height []int) int {
	sum := 0
	for {
		s := -1
		for i, v := range height {
			if v > 0 {
				if s != -1 {
					sum += i - s - 1
				}
				s = i
				height[i]--
			}
		}
		if s == -1 {
			break
		}
	}
	return sum
}

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	sum := 0
	left := 0
	left_max := height[left]
	right := len(height) - 1
	right_max := height[right]
	for left != right {
		if height[right] > height[left] {
			if left_max < height[left] {
				left_max = height[left]
			}
			sum += left_max - height[left]
			left++
		} else {
			if right_max < height[right] {
				right_max = height[right]
			}
			sum += right_max - height[right]
			right--
		}
	}
	return sum
}
