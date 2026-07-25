func rob(nums []int) int {	
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	helper := func (nums []int) int{
		prev, cur := 0, 0

		for _, num := range nums{
				prev, cur = cur, max(prev+num, cur)
		}
		return cur
		}
		return max(helper(nums[1:]),helper(nums[:len(nums)-1]))

}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
