package tree

/*
https://leetcode.cn/problems/maximum-depth-of-binary-tree/
算出左右子树的最大深度然后加一

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return GetMax(root)
}
func GetMax(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftNum := GetMax(root.Left)
	rightNum := GetMax(root.Right)
	if leftNum > rightNum {
		return 1 + leftNum
	} else {
		return 1 + rightNum
	}
}
