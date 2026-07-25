/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {


	visited := make(map[*Node]*Node)

	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node{
		if node == nil {
			return nil
		}
		if _, found := visited[node]; found {
			return visited[node]
		}

		copy := &Node{Val: node.Val}
		visited[node] = copy
		for _, nei := range node.Neighbors{
			copy.Neighbors = append(copy.Neighbors, dfs(nei))
		}
		return copy
	}

	return dfs(node)
}
