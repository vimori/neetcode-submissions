/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	curIndex := 0
	res := 0
	var dfs func(*TreeNode) 
	dfs = func(root *TreeNode) {
		if root == nil{
			return 
		}
		dfs(root.Left)
		curIndex++
		if curIndex == k{
			res = root.Val
			return 
		}
		dfs(root.Right)
		return 
	}

	dfs(root)
	return res
}
