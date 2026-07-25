/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    var isSameTree func(*TreeNode, *TreeNode) bool
	isSameTree = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil{
			return true
		}
		if (p != nil && q == nil) || (p == nil && q != nil){
			return false
		}
		if p.Val != q.Val {
			return false
		}

		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	var hasSubtree func(*TreeNode) bool
	hasSubtree = func(root *TreeNode) bool{
		if root == nil {
			return false
		}
		if isSameTree(root, subRoot) {
			return true
		}

		return hasSubtree(root.Left) || hasSubtree(root.Right)
	}

	return hasSubtree(root)
}
