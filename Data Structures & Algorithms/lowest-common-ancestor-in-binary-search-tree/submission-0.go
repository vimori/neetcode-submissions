/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	lca := root

	var search func(*TreeNode)
	search = func(root *TreeNode){
		if root == nil{
			return 
		}

		lca = root
		if root == p || root == q{
			return 
		} else if p.Val > root.Val && q.Val > root.Val{
			search(root.Right)
		} else if p.Val < root.Val && q.Val < root.Val{
			search(root.Left)
		}else{
			return
		}
	}
	search(root)
	return lca
}
