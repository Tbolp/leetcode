package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := (len(nums) - 1) / 2
	res := &TreeNode{Val: nums[root]}
	res.Left = sortedArrayToBST(nums[0:root])
	res.Right = sortedArrayToBST(nums[root+1:])
	return res
}
