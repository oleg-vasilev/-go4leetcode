package lc541

// #542 - 01 Matrix

// Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
// The distance between two adjacent cells is 1.

// Constraints:
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 104
// 1 <= m * n <= 104
// mat[i][j] is either 0 or 1.
// There is at least one 0 in mat.

var directionalShifts = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func updateMatrix(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])
	var queue [][2]int
	
	// preprocess matrix
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// add zeros to queue and mark cells that should be calculated with big number
			// obviously greater than the possible distance value
			if mat[row][col] == 0 {
				queue = append(queue, [2]int{row, col})
			} else {
				mat[row][col] = 99999
			}
		}
	}
	// process queue
	for len(queue) > 0 {
		// dequeue next item
		current := queue[0]
		queue = queue[1:]
		// check all 4 neighbours
		for _, shift := range directionalShifts {
			x, y := current[0]+shift[0], current[1]+shift[1]
			// skip impossible and solved neighbours
			if x < 0 || x >= rows ||
				y < 0 || y >= cols ||
				mat[x][y] <= mat[current[0]][current[1]] {
				continue
			}
			// enqueue neighbour
			queue = append(queue, [2]int{x, y})
			// increase distance value in the cell
			mat[x][y] = mat[current[0]][current[1]] + 1
		}
	}
	
	return mat
}
