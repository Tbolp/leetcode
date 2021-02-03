package numarray

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	obj := NumArray{}
	sum := 0
	obj.nums = make([]int, len(nums)+1)
	for i, v := range nums {
		obj.nums[i] = sum
		sum += v
	}
	obj.nums[len(nums)] = sum
	return obj
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.nums[j+1] - this.nums[i]
}
