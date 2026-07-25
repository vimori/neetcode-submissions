/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    balanced := true 

	var dfs func (*TreeNode) int
	dfs = func (root *TreeNode) int {
		if root == nil{
			return 0
		}

		left_height := dfs(root.Left)
		if !balanced {
			return 0
		}
		right_height := dfs(root.Right)

		if abs(left_height-right_height) > 1 {
    		balanced = false
		}
		return 1 + max(left_height,right_height)
	}
	dfs(root)
	return balanced
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
