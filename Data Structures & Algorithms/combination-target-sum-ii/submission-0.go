func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	nums := candidates
	sort.Ints(nums)

	var backtracking func(i int, cur []int, total int) 
	backtracking = func(i int, cur []int, total int) {
		if total == target{
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		if i >= len(nums) || total > target {
			return
		}
		cur = append(cur, nums[i])
		backtracking(i+1, cur, total + nums[i])
		cur = cur[:len(cur)-1]
		for i+1 < len(nums) &&  nums[i] == nums[i+1]{
			i++
		}
		backtracking(i+1, cur, total)
	}

	backtracking(0, []int{}, 0)
	return res
}
