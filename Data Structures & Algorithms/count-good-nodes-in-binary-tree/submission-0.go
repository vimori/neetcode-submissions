/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	if root == nil {
        return 0
    }

	res := 0

	q := []struct {
		node *TreeNode
		maxVal int
	}{{root, math.MinInt32}}

	for len(q) > 0{
		front := q[0]
		q = q[1:]

		node := front.node
		maxVal := front.maxVal

		if node.Val >= maxVal {
			res++
		}

		newMaxVal := max(maxVal, node.Val)

		if node.Left != nil{
			q = append(q, struct{
				node *TreeNode
				maxVal int
			}{node.Left, newMaxVal})
		}

		if node.Right != nil{
			q = append(q, struct{
				node *TreeNode
				maxVal int
			}{node.Right, newMaxVal})
		}

	}

	return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}