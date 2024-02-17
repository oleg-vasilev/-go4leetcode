package lc695

// 695. Max Area of Island

// You are given an m x n binary matrix grid. An island is a group of
// 1's (representing land) connected 4-directionally (horizontal or vertical.)
// You may assume all four edges of the grid are surrounded by water.
//
// The area of an island is the number of cells with a value 1 in the island.
//
// Return the maximum area of an island in grid. If there is no island, return 0.

// TODO : Use original grid to mark visits to save memory

// func maxAreaOfIsland(grid [][]int) int {
// 	// sooner or later we will visit all items so to track visits array with same dimensions
// 	visited := make([][]bool, len(grid))
// 	for ind, row := range grid {
// 		visited[ind] = make([]bool, len(row))
// 	}
// 	var maxSize, currentSize int
// 	for row := range grid {
// 		for col := range grid[row] {
// 			if grid[row][col] == 1 && !visited[row][col] {
// 				// we find an island - start measuring
// 				currentSize = dfs(grid, visited, row, col)
// 				if currentSize > maxSize {
// 					maxSize = currentSize
// 				}
// 			}
// 		}
// 	}
// 	return maxSize
// }

// func dfs(grid [][]int, visited [][]bool, row, col int) int {
// 	visited[row][col] = true
// 	weight := 1
// 	// left
// 	if col-1 >= 0 && grid[row][col-1] == 1 && !visited[row][col-1] {
// 		weight += dfs(grid, visited, row, col-1)
// 	}
// 	// right
// 	if col+1 < len(grid[row]) && grid[row][col+1] == 1 && !visited[row][col+1] {
// 		weight += dfs(grid, visited, row, col+1)
// 	}
// 	// up
// 	if row-1 >= 0 && grid[row-1][col] == 1 && !visited[row-1][col] {
// 		weight += dfs(grid, visited, row-1, col)
// 	}
// 	// down
// 	if row+1 < len(grid) && grid[row+1][col] == 1 && !visited[row+1][col] {
// 		weight += dfs(grid, visited, row+1, col)
// 	}
// 	return weight
// }

// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: object-oriented solution "for-fun"

// func maxAreaOfIsland(grid [][]int) int {
// 	var visited visitsMap = make(map[gridCell]struct{})
// 	var maxSize, currentSize int
// 	for row := range grid {
// 		for col := range grid[row] {
// 			cell := gridCell{row, col}
// 			if cell.isLand(grid) && !visited.contains(cell) {
// 				currentSize = visitCell(grid, visited, cell)
// 				if currentSize > maxSize {
// 					maxSize = currentSize
// 				}
// 			}
// 		}
// 	}
// 	return maxSize
// }

// func visitCell(grid [][]int, visited visitsMap, currentCell gridCell) int {
// 	visited[currentCell] = struct{}{}
// 	weight := 1
// 	// check left
// 	leftCell := currentCell.left()
// 	if leftCell.isValid(grid) && leftCell.isLand(grid) && !visited.contains(leftCell) {
// 		weight += visitCell(grid, visited, leftCell)
// 	}
// 	// check right
// 	rightCell := currentCell.right()
// 	if rightCell.isValid(grid) && rightCell.isLand(grid) && !visited.contains(rightCell) {
// 		weight += visitCell(grid, visited, rightCell)
// 	}
// 	// check top
// 	topCell := currentCell.top()
// 	if topCell.isValid(grid) && topCell.isLand(grid) && !visited.contains(topCell) {
// 		weight += visitCell(grid, visited, topCell)
// 	}
// 	// check bottom
// 	bottomCell := currentCell.bottom()
// 	if bottomCell.isValid(grid) && bottomCell.isLand(grid) && !visited.contains(bottomCell) {
// 		weight += visitCell(grid, visited, bottomCell)
// 	}
// 	return weight
// }

// type gridCell struct {
// 	row, col int
// }

// func (c gridCell) left() gridCell {
// 	return gridCell{row: c.row, col: c.col - 1}
// }
// func (c gridCell) right() gridCell {
// 	return gridCell{row: c.row, col: c.col + 1}
// }
// func (c gridCell) top() gridCell {
// 	return gridCell{row: c.row - 1, col: c.col}
// }
// func (c gridCell) bottom() gridCell {
// 	return gridCell{row: c.row + 1, col: c.col}
// }
// func (c gridCell) isValid(grid [][]int) bool {
// 	if c.col < 0 || c.row < 0 {
// 		return false
// 	}
// 	if c.row >= len(grid) || c.col >= len(grid[c.row]) {
// 		return false
// 	}
// 	return true
// }
// func (c gridCell) isLand(grid [][]int) bool {
// 	return grid[c.row][c.col] == 1
// }

// type visitsMap map[gridCell]struct{}

// func (vc visitsMap) contains(c gridCell) bool {
// 	_, ok := vc[c]
// 	return ok
// }

// :::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::: iterative_solution

func maxAreaOfIsland(grid [][]int) int {
	visited := map[[2]int]struct{}{}
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	area := 0
	for i, row := range grid {
		for j, _ := range row {
			stack := [][2]int{{i, j}}
			currArea := 0
			for len(stack) > 0 {
				pair := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				r := pair[0]
				c := pair[1]
				// skip 'water' cells filled with 0
				if grid[r][c] == 0 {
					continue
				}
				// also skip visited cells
				if _, ok := visited[pair]; ok {
					continue
				}
				visited[pair] = struct{}{}
				currArea++
				for _, move := range directions {
					row := r + move[0]
					col := c + move[1]
					if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
						// skip impossible cells
						continue
					}
					stack = append(stack, [2]int{row, col})
				}
			}
			if currArea > area {
				area = currArea
			}
			
		}
	}
	return area
}
