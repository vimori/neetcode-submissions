func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		a := nums[i]
		if nums[i] > 0 {
			break
		}
		if i > 0 && a == nums[i-1]{
			continue
		}

		target := 0 - nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			curSum := nums[l] + nums[r]
			if curSum < target{
				l++
			} else if curSum > target{
				r--
			} else{
				result = append(result, []int{a, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1]{
					l++
				}
			}
		}
	}
	return result
}
