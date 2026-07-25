func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}

    var backtracking func(int, []int, int)
    backtracking = func(i int, cur []int, total int) {
		if total == target {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		if i >= len(nums) || total > target{
			return
		}
		cur = append(cur, nums[i])
		backtracking(i, cur, total + nums[i])
		cur = cur[:len(cur)-1]
		backtracking(i+1, cur, total) // не берём nums[i](2 случай ветки)
	}

	backtracking(0, []int{}, 0)
	return res
}
