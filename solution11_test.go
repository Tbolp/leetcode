package leetcode

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	area := func(i, j int) int {
		if height[i] < height[j] {
			return (j - i) * height[i]
		} else {
			return (j - i) * height[j]
		}
	}
	max := 0
	for end > start {
		tmp := area(start, end)
		if tmp > max {
			max = tmp
		}
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return max
}
