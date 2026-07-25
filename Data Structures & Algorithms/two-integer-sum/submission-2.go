func twoSum(nums []int, target int) []int {
    indices := make(map[int]int)

    for i,num := range nums{
        indices[num] = i
    }

    for i,num := range nums{
        diff := target - num
        if j, found := indices[diff]; found && j != i{
            return []int{i,j}
        }
    }
    return []int{}
}
