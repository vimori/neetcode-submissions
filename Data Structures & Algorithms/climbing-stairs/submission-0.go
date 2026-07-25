func climbStairs(n int) int {
    one := 1
	two := 1

	for i:= 0; i < n-1; i++{
		temp := one
		one += two
		two = temp
	}

	return one
}
