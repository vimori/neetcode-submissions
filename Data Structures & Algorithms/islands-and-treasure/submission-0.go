func islandsAndTreasure(grid [][]int) {
    if len(grid) == 0 {
		return 
 	}
	const INF = 2147483647
	rows, cols := len(grid), len(grid[0])
	queue := [][]int{}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++{
			if grid[i][j] == 0{
				queue = append(queue, []int{i,j})
			}
		} 
	}

	dirs := [][]int{
		{0,1},
		{0,-1},
		{1, 0},
		{-1,0},
	}

	for len(queue) > 0 {
		cell := queue[0]
		i, j := cell[0], cell[1]
		queue = queue[1:]

		for _, dir := range dirs{
			ni, nj := i+dir[0], j+dir[1]	

			if ni >= 0 && ni < rows && nj >= 0 && 
			nj < cols && grid[ni][nj] == INF {
				grid[ni][nj] = grid[i][j] + 1
				queue = append(queue, []int{ni, nj})
			}
		}

	}
}
