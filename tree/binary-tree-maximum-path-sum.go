package tree

import "strconv"

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/815690/yi-pian-wen-zhang-jie-jue-suo-you-er-cha-kcb0/
// 路径题汇总

// 路径类题分为自顶向下和非自顶向下
// 求从根节点到叶子节点的所有路径https://leetcode.cn/problems/binary-tree-paths/description/
// 如果只有根节点，就直接返回根节点，否则，依次把左右子树的路径拼接到根节点上
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, strconv.Itoa(root.Val))
	}
	leftVal := binaryTreePaths(root.Left)
	rightVal := binaryTreePaths(root.Right)
	for _, val := range leftVal {
		res = append(res, strconv.Itoa(root.Val)+"->"+val)
	}
	for _, val := range rightVal {
		res = append(res, strconv.Itoa(root.Val)+"->"+val)
	}
	return res
}

// 是否从根到叶子路径和等于sum的路径https://leetcode.cn/problems/path-sum/description/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 获取从根到叶子路径和等于sum的路径https://leetcode.cn/problems/path-sum-ii/
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		temp := make([]int, 0)
		temp = append(temp, root.Val)
		res = append(res, temp)
		return res
	}
	leftVal := pathSum(root.Left, targetSum-root.Val)
	rightVal := pathSum(root.Right, targetSum-root.Val)
	for _, val := range leftVal {
		res = append(res, append([]int{root.Val}, val...))
	}
	for _, val := range rightVal {
		res = append(res, append([]int{root.Val}, val...))
	}
	return res
}

// 路径总和，https://leetcode.cn/problems/path-sum-iii/
// 不从根节点开始，双重递归，不从叶子结束，递归结束不需要判断左右为空，
func PathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
    // 先以根节点算一遍
	res += HasPathSum(root, targetSum)
    // 再以左右算一遍
	res += PathSum(root.Left, targetSum)
	res += PathSum(root.Right, targetSum)
	return res
}
func HasPathSum(root *TreeNode, targetSum int) int {
	res := 0
	if root == nil {
		return 0
	}
	if root.Val == targetSum {
		res += 1
	}
	res += HasPathSum(root.Left, targetSum-root.Val)
	res += HasPathSum(root.Right, targetSum-root.Val)
	return res
}
