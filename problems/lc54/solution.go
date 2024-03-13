package lc54

// 54. Spiral Matrix

const (
	right = iota
	down
	left
	up
)

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	columns := len(matrix[0])
	result := make([]int, rows*columns)

	var row, col int
	var counter int
	var direction = right

	for counter < rows*columns {
		result[counter] = matrix[row][col]
		switch direction {
		case right:
			if row+col == columns-1 {
				direction = down
				row++
			} else {
				col++
			}
		case down:
			if columns-col == rows-row {
				direction = left
				col--
			} else {
				row++
			}
		case left:
			if row+col == rows-1 {
				direction = up
				row--
			} else {
				col--
			}
		case up:
			if row-col == 1 {
				direction = right
				col++
			} else {
				row--
			}
		}
		counter++
	}
	return result
}
