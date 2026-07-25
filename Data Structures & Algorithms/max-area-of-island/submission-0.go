func maxAreaOfIsland(grid [][]int) int {
    directions := [][]int{
		{1,0},
		{-1,0},
		{0,1},
		{0,-1},
	}
	rows, cols := len(grid), len(grid[0])
	maxArea := 0

	var bfs func(r,c int) int 
	bfs = func(r,c int) int{
		curArea := 1
		q := [][]int{{r,c}}
		grid[r][c] = 0

		for len(q) > 0 {
			node := q[0]
			q = q[1:]
			row, col := node[0], node[1]
			for _, dir := range directions{
				nr, nc := row + dir[0], col + dir[1]
				if nr < 0 || nr >= rows || nc < 0 ||
				nc >= cols || grid[nr][nc] == 0{
					continue
				}
				q = append(q, []int{nr,nc})
				grid[nr][nc] = 0
				curArea++
			} 
		}
		return curArea
	}	

	for r := 0; r < rows; r++{
		for c := 0; c < cols; c++{
			if grid[r][c] == 1 {
				curArea := bfs(r,c)
				if curArea > maxArea{
					maxArea = curArea
				}
			}
		}
	}
	return maxArea
}
