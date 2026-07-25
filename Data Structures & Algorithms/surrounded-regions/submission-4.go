func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	var dfs func(r, c int)
	dfs = func(r, c int){
		if r < 0 || c < 0 || r >= m || c >= n || board[r][c] != 'O' {
			return
		}
		board[r][c] = 'T'
		dfs(r+1,c)
		dfs(r-1,c)
		dfs(r,c+1)
		dfs(r,c-1)
	}

	for r := 0; r < m; r++{
		dfs(r,0)
		dfs(r,n-1)
	}

	for c := 0; c < n; c++{
		dfs(0,c)
		dfs(m-1,c)
	}

	for r := 0; r < m; r++{
		for c := 0; c < n; c++{
			if board[r][c] == 'O'{
				board[r][c] = 'X'
			}
		}
	}
	
	for r := 0; r < m; r++{
		for c := 0; c < n; c++{
			if board[r][c] == 'T'{
				board[r][c] = 'O'
			}
		}
	}

}
