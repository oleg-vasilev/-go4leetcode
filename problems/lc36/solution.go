package lc36

// 36. Valid Sudoku

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:
//
// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

// Constraints:
//
// board.length == 9
// board[i].length == 9
// board[i][j] is a digit 1-9 or '.'.

//  constant space ~O(1) and linear time ~O(m+n+mn) solution

func isValidSudoku(board [][]byte) bool {
	var validator = Validator{}
	// validate rows
	for row := range board {
		validator.Clear() // clear validator before each row
		for col := range board[row] {
			if ok := validator.CheckValue(board[row][col]); !ok {
				return false
			}
		}
	}
	// validate columns
	for col := 0; col < len(board); col++ {
		validator.Clear() // clear validator before each row
		for row := 0; row < len(board); row++ {
			if ok := validator.CheckValue(board[row][col]); !ok {
				return false
			}
		}
	}

	// validate squares

	// Coords range cheatsheet for each square (row,col)
	// 0th:(0-2,0-2) 1st:(0-2,3-5), 2nd:(0-2,6-8)
	// 3rd:(3-5,0-2) 4th:(3-5,3-5), 5th:(3-5,6-8)
	// 6th:(6-8,0-2) 7th:(6-8,3-5), 8th:(6-8,6-8)
	var row, col int // temp values to calculate cell coords
	for square := 0; square < 9; square++ {
		validator.Clear() // clear validator before each square
		for cell := 0; cell < 9; cell++ {
			// local cell coords in 3x3 grid based on item number
			row, col = cell/3, cell%3
			// adjust local coords to global 9x9 coords
			row += 3 * (square / 3)
			col += 3 * (square % 3)
			// add cell to validator
			if ok := validator.CheckValue(board[row][col]); !ok {
				return false
			}
		}
	}
	return true
}

// Validator
// 1-9th items used as value presence flags,
// 0th - as invalid flag (don't needed actually)
type Validator [10]bool

func (v *Validator) Clear() {
	for i := range v {
		v[i] = false
	}
}

func (v *Validator) Invalid() bool {
	return v[0]
}

func (v *Validator) isPossibleValue(val int) bool {
	if val > 9 || val < 1 || v[val] {
		// invalid value or value already exist
		v[0] = true  // set invalid flag
		return false // false - value is invalid
	}
	v[val] = true
	return true // value is valid
}

func (v *Validator) CheckValue(val byte) bool {
	cellValue := int(val)
	// '.' is 46, 1-9 is 49-57
	if cellValue == 46 {
		// empty cell - return ok
		return true
	} else if cellValue >= 49 && cellValue <= 57 {
		// valid number - check
		return v.isPossibleValue(cellValue - 48)
	} else {
		return false // just invalid input
	}
}
