func solveNQueens(n int) [][]string {
	res := [][]string{}
	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	col := make(map[int]bool)
	posDiag := make(map[int]bool)
	negDiag := make(map[int]bool)

	var backtrack func(r int)
    backtrack = func(r int) {	
		if r == n{
			solution := make([]string, n)
			for i := range board {
				solution[i] = string(board[i])
			}
			res = append(res, solution)
			return
		}

		for c := 0; c < n; c++ {
			if col[c] || posDiag[r+c] || negDiag[r-c] {
				continue
			}

			col[c] = true
			posDiag[r+c] = true
			negDiag[r-c] = true
			board[r][c] = 'Q'

			backtrack(r + 1)

			col[c] = false
			posDiag[r+c] = false
			negDiag[r-c] = false
			board[r][c] = '.'
		}
	}	
	backtrack(0)
	return res
}
