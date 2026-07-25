func orangesRotting(grid [][]int) int {
    if len(grid) == 0{
		return 0
	}
	minutes := 0
	rows, cols := len(grid), len(grid[0])
	queue := [][]int{}
	fresh := 0

	for i := 0; i < rows; i++{
		for j := 0; j < cols; j++{
			if grid[i][j] == 2{
				queue = append(queue, []int{i,j})
			}else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	if fresh == 0 {
		return 0
	}

	dirs := [][]int{
		{1,0},
		{-1,0},
		{0,1},
		{0,-1},
	}
	for len(queue) > 0{
		levelSize := len(queue)
		for i := 0; i < levelSize; i++{
			cell := queue[0]
			queue = queue[1:]
			r, c := cell[0],cell[1]

			for _, dir := range dirs{
				nr, nc := r + dir[0], c + dir[1]
				if nc >= 0 && nc < rows && nr >= 0 && 
				nr < cols && grid[nc][nr] == 1{
					grid[nc][nr] = 2
					fresh--
					queue = append(queue, []int{nc,nr})
				}
			}
		}
		minutes++
	}

	if fresh > 0 {
		return -1
	}

	return minutes - 1 
}

