/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type QueueItem struct {
	node *TreeNode
	min int64
	max int64
}

func isValidBST(root *TreeNode) bool {
    q := []QueueItem{{root, math.MinInt64, math.MaxInt64}}

    for len(q) > 0 {
        item := q[0]
		q = q[1:]

		val := int64(item.node.Val)
		if val <= item.min || val >= item.max {
			return false
        }

        if item.node.Left != nil {
            q = append(q, QueueItem{item.node.Left, item.min, val}) // min, max
        }
        if item.node.Right != nil {
            q = append(q, QueueItem{item.node.Right, val, item.max}) // min, max
        }

	}
    return true
}
