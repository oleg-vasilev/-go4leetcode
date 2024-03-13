package lc48

import (
	"slices"
)

// 48. Rotate Image

// You are given an n x n 2D matrix representing an image,
// rotate the image by 90 degrees (clockwise).
//
// You have to rotate the image in-place, which means you have
// to modify the input 2D matrix directly.
// DO NOT allocate another 2D matrix and do the rotation.
//

func rotate(matrix [][]int) {
	// so simple rules are :
	// Nth in row becomes Nth in column
	// Nth in column becomes Size-n in row
	// or
	// transpose then reverse rows

	// transposing square matrix inplace
	for n := 0; n < len(matrix); n++ {
		for idx := n + 1; idx < len(matrix); idx++ {
			matrix[n][idx], matrix[idx][n] = matrix[idx][n], matrix[n][idx]
		}
	}
	// reversing rows
	for i := range matrix {
		slices.Reverse(matrix[i])
	}
}
