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

	var col, row int
	var counter int
	var direction = right

	for counter < rows*columns {
		result[counter] = matrix[col][row]
		switch direction {
		case right:
			if col+row == columns-1 {
				direction = down
				col++
			} else {
				row++
			}
		case down:
			if columns-row == rows-col {
				direction = left
				row--
			} else {
				col++
			}
		case left:
			if col+row == rows-1 {
				direction = up
				col--
			} else {
				row--
			}
		case up:
			if col-row == 1 {
				direction = right
				row++
			} else {
				col--
			}
		}
		counter++
	}
	return result
}
