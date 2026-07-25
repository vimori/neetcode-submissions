/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    res := []int{}
    q := []*TreeNode{root}

    for len(q) > 0 {
        rightSide := 0
        qLen := len(q)

        for i := 0; i < qLen; i++ {
            node := q[0]
            q = q[1:]

            if node != nil {
                rightSide = node.Val
                if node.Left != nil {
                    q = append(q, node.Left)
                }
                if node.Right != nil {
                    q = append(q, node.Right)
                }
            }
        }
        res = append(res, rightSide)
    }

    return res
}