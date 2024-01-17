package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func createTree(nodes []int) *TreeNode {
    if len(nodes) == 0 {
        return nil
    }

    root := &TreeNode{Val: nodes[0]}
    stack := []*TreeNode{root}

    for i := 1; i < len(nodes); i++ {
        node := stack[len(stack)-1]

        if node.Left == nil {
            newNode := &TreeNode{Val: nodes[i]}
            node.Left = newNode
            stack = append(stack, newNode)
        } else if node.Right == nil {
            newNode := &TreeNode{Val: nodes[i]}
            node.Right = newNode
            stack = append(stack, newNode)
        } else {
            stack = stack[:len(stack)-1]
            i--
        }
    }

    return root
}
func main() {
	nodes := []int{0, 1, 2, 3, 4, 3, 4}
	root := createTree(nodes)
    smallestFromLeaf(root)
}
func smallestFromLeaf(root *TreeNode) string {
    res := make([]string, 0)
    dfs(root, "", res)
    return res[0]
}
func dfs(root *TreeNode, s string, res []string){
    if root == nil{
        return
    }
    if root.Left == nil && root.Right == nil{
        s = s + string('a' + root.Val)
        res = append(res, s)
        return 
    }
    dfs(root.Left, s, res)
    dfs(root.Right, s, res)
}
func reverse(s string) (res string){
    for _, c := range s{
        res = string(c) + res
    }
    return res
}
