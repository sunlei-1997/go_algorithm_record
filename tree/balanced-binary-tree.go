package tree

/**
https://leetcode.cn/problems/balanced-binary-tree/
函数判断是否是平衡二叉树，只有左子树和右子树都是，且左子树高度减去右子树高度的绝对值小于等于1
 */
func isBalanced(root *TreeNode) bool {
	return cal(root)
}
func cal(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHe := GetMax(root.Left)
	rightHe := GetMax(root.Right)

	return cal(root.Left) && cal(root.Right) && (leftHe-rightHe >= -1) && (leftHe-rightHe <= 1)
}
