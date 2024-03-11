package lc6

import (
	"strings"
)

// 6. Zigzag Conversion

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given
// number of rows like this: (you may want to display
// this pattern in a fixed font for better legibility)
//
// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"
//
// Write the code that will take a string and make this conversion given a number of rows:
//
// string convert(string s, int numRows);

// Actual simulation
func convertSimulation(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	var row, col int
	var res = make([][]string, numRows)
	for i, _ := range res {
		res[i] = make([]string, len(s))
	}

	var maxVerticalItems = numRows
	var maxDiagonalItems = numRows - 2

	var currVerticalLeft = maxVerticalItems
	var currDiagLeft = maxDiagonalItems

	for _, v := range s {
		// logic for general "snake" fill down-up-down-up-etc.
		// {
		// 	// find a place for letter
		// 	col = i / numRows
		// 	if col%2 == 0 {
		// 		row = i % numRows
		// 	} else {
		// 		// reverse order - bottom to up
		// 		row = (numRows - 1) - (i % numRows)
		// 		// additional column shift
		// 	}
		// }

		// fill the item
		res[row][col] = string(v)
		// move coords for next iteration
		if currVerticalLeft > 0 {
			// placed vertical item on same column but lower row
			currVerticalLeft--
			if currVerticalLeft == 0 {
				row--
				col++
			} else {
				row++
			}
			// continue
		} else if currDiagLeft > 0 {
			// place diagonal item at next column and previous row
			currDiagLeft--
			row--
			col++
		}
		if currVerticalLeft == 0 && currDiagLeft == 0 {
			currVerticalLeft = maxVerticalItems
			currDiagLeft = maxDiagonalItems
		}
	}
	// convert array to string in left-to-right-up-to-down order
	var builder = strings.Builder{}
	builder.Grow(len(s))
	for _, v := range res {
		builder.WriteString(strings.Join(v, ""))
	}
	return builder.String()
}

func convert(s string, numRows int) string {
	// basic cases
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	var rows = make([]string, numRows)
	var curr = 0       // current row - initially we ge from top to bottom
	var direction bool // false - down, true - up

	for _, v := range s {
		// place char to current row
		rows[curr] += string(v)
		// change row for next iteration according to current direction
		if direction {
			curr--
		} else {
			curr++
		}
		// if we reach bound - invert direction for next iteration
		if curr == 0 || curr == numRows-1 {
			direction = !direction
		}
	}
	return strings.Join(rows, "")
}
