package lc200

// 200. Number of Islands

// Given an m x n 2D binary grid grid which represents a map
// of '1's (land) and '0's (water), return the number of islands.
//
// An island is surrounded by water and is formed by connecting
// adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.
//

func numIslands(grid [][]byte) int {
	var count int
	// idea is to traverse the field until reach land ('1')
	// and somehow find and mark all adjacent land cells as traversed
	// to skip them in future iterations
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '1' {
				// found a land cell
				count++
				// iterate over all neighbour land cells and mark them
				islandMarkDFSRecursive(grid, row, col)
				// islandMarkBFSIterative(grid, row, col)
			}
		}
	}
	return count
}

func isValidLandCell(grid [][]byte, row, col int) bool {
	return row >= 0 && row < len(grid) &&
		col >= 0 && col < len(grid[row]) &&
		grid[row][col] == '1'
}

func islandMarkDFSRecursive(grid [][]byte, row, col int) {
	// invalid input - out of bounds or not an land cell
	// if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) || grid[row][col] != '1' {
	if !isValidLandCell(grid, row, col) {
		return
	}
	// mark current cell
	grid[row][col] = 'x'
	// run DFS on neighbours
	islandMarkDFSRecursive(grid, row+1, col)
	islandMarkDFSRecursive(grid, row-1, col)
	islandMarkDFSRecursive(grid, row, col+1)
	islandMarkDFSRecursive(grid, row, col-1)
}

func islandMarkBFSIterative(grid [][]byte, row, col int) {
	var neighbourOffsets = [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	// construct a queue with only initial cell
	var queue [][2]int
	queue = append(queue, [2]int{row, col})
	// process cells while queue isn't empty
	for len(queue) > 0 {
		// dequeue
		coords := queue[0]
		queue = queue[1:]
		// mark as visited
		grid[coords[0]][coords[1]] = 'x'
		// iterate over neighbours
		for _, offset := range neighbourOffsets {
			nRow := coords[0] + offset[0]
			nCol := coords[1] + offset[1]
			// if neighbour is valid land - add it to queue
			if isValidLandCell(grid, nRow, nCol) {
				queue = append(queue, [2]int{nRow, nCol})
			}
		}
	}
}
