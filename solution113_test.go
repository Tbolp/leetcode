package leetcode

var pathSum_Context struct {
	Res [][]int
	Tmp []int
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	pathSum_Context.Res = [][]int{}
	pathSum_Context.Tmp = []int{}
	pathSum_(root, targetSum)
	return pathSum_Context.Res
}

func pathSum_(root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			tmp := make([]int, len(pathSum_Context.Tmp)+1)
			copy(tmp, pathSum_Context.Tmp)
			tmp[len(tmp)-1] = targetSum
			pathSum_Context.Res = append(pathSum_Context.Res, tmp)
		}
	} else {
		pathSum_Context.Tmp = append(pathSum_Context.Tmp, root.Val)
		pathSum_(root.Left, targetSum-root.Val)
		pathSum_(root.Right, targetSum-root.Val)
		pathSum_Context.Tmp = pathSum_Context.Tmp[:len(pathSum_Context.Tmp)-1]
	}
}
