func pacificAtlantic(heights [][]int) [][]int {
    if len(heights) == 0 {
		return [][]int{}
	}

	rows, cols := len(heights), len(heights[0])

	pac := make([][]bool, rows)
	atl := make([][]bool, rows)
	for i := 0; i < rows; i++{
		pac[i] = make([]bool, cols)
		atl[i] = make([]bool, cols)
	}

	pacStart := [][2]int{}
	atlStart := [][2]int{}

	for r := 0; r < rows; r++{
		pacStart = append(pacStart, [2]int{r,0})
		atlStart = append(atlStart, [2]int{r,cols-1})

	}
	for c := 0; c < cols; c++{
		pacStart = append(pacStart, [2]int{0,c})
		atlStart = append(atlStart, [2]int{rows-1,c})
	}

	dirs := [][]int{
		{1,0},
		{-1,0},
		{0,1},
		{0,-1},
	}

	bfs := func(queue [][2]int, ocean [][]bool) {
		for len(queue) > 0{
		cell := queue[0]
		queue = queue[1:]
		r, c := cell[0], cell[1] 
		ocean[r][c] = true

		for _, dir := range dirs{
			nr, nc := r + dir[0], c +  dir[1]
			if nc >= 0 && nc < cols && nr >= 0 && nr < rows &&
				 !ocean[nr][nc] && heights[nr][nc] >= heights[r][c]{
					queue = append(queue,[2]int{nr,nc})
				}
		}
	}
	}
	bfs(pacStart, pac)
	bfs(atlStart, atl)
	result := [][]int{}
	for r := 0; r < rows; r++{
		for c := 0; c < cols; c++{
			if pac[r][c] && atl[r][c]{
				result = append(result, []int{r,c})
			}
		}
	}
	return result
}
