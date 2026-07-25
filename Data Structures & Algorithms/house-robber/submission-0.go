func rob(nums []int) int {
    prev, cur := 0, 0

	for _, num := range nums{
		prev, cur = cur, max(num+prev, cur)
	}
	return cur
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}