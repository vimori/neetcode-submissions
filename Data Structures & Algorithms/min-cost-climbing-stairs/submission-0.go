func minCostClimbingStairs(cost []int) int {
    n := len(cost)

	for i := n - 3; i >= 0; i--{
		cost[i] += min(cost[i+1], cost[i+2])
	}
	return min(cost[0], cost[1])
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}