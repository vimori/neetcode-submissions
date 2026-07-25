func permute(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}

	var backtracking func()
	backtracking = func(){
		if len(cur) == len(nums){
			temp := make([]int, len(nums))
			copy(temp, cur)
			res = append(res, temp)
		}
		for _, val := range nums{
			if !contains(val, cur){
			cur = append(cur, val)
			backtracking()
			cur = cur[:len(cur)-1]
			}

		}
	}
	backtracking()
	return res
}

func contains(target int, slice []int) bool {
    for _, v := range slice {
        if v == target {
            return true
        }
    }
    return false
}