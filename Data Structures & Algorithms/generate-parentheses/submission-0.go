func generateParenthesis(n int) []string {
	res := []string{}
	cur := []string{}

	var backtracking func(open int, close int)
	backtracking = func(open int, close int){
		if len(cur) == 2*n{
			res = append(res, strings.Join(cur, ""))
			return 
		}
		if open < n{
			cur = append(cur, "(")
			backtracking(open+1, close)
			cur = cur[:len(cur)-1]
		}
		if close < open {
			cur = append(cur, ")")
			backtracking(open, close+1)
			cur = cur[:len(cur)-1]
		}
	}
	backtracking(0,0)
	return res
}
