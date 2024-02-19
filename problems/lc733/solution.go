package floodfill

// 733. Flood Fill

// An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.
//
// You are also given three integers sr, sc, and color. You should perform a flood fill on the image starting from the pixel image[sr][sc].
//
// To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with color.
//
// Return the modified image after performing the flood fill.

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	// result := make([][]int, len(image)) // to save the memory
	// copy(result, image)
	visited := make(map[coords]struct{})
	dfs(image, visited, coords{sr, sc}, color)
	return image
}

type coords struct {
	row, col int
}

func dfs(field [][]int, visited map[coords]struct{}, cell coords, fill int) {
	if _, ok := visited[cell]; ok {
		return
	}
	visited[cell] = struct{}{}
	originalColor := field[cell.row][cell.col]
	field[cell.row][cell.col] = fill
	// left
	if cell.col-1 >= 0 && field[cell.row][cell.col-1] == originalColor {
		dfs(field, visited, coords{col: cell.col - 1, row: cell.row}, fill)
	}
	// right
	if cell.col+1 < len(field[cell.row]) && field[cell.row][cell.col+1] == originalColor {
		dfs(field, visited, coords{col: cell.col + 1, row: cell.row}, fill)
	}
	// up
	if cell.row-1 >= 0 && field[cell.row-1][cell.col] == originalColor {
		dfs(field, visited, coords{col: cell.col, row: cell.row - 1}, fill)
	}
	// down
	if cell.row+1 < len(field) && field[cell.row+1][cell.col] == originalColor {
		dfs(field, visited, coords{col: cell.col, row: cell.row + 1}, fill)
	}
}

// TODO: implement non-recursive algorithm
