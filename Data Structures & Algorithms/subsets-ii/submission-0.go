func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}
	sort.Ints(nums)

	var backtracking func(int)
	backtracking = func(i int){
		if i == len(nums){
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return 
		}

		cur = append(cur, nums[i])
		backtracking(i+1)
		cur = cur[:len(cur)-1]
		for i+1 < len(nums) && nums[i] == nums[i+1]{
			i++
		}
		backtracking(i+1)

	}
	backtracking(0)
	return res
}
