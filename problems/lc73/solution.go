package lc73

// 73. Set Matrix Zeroes

// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

// You must do it in place.

// in place solution with O(1) extra space and O(mn) time.

func setZeroes(matrix [][]int) {
	// main trick is to distinct 'real' 0 in head row and column from 0s placed here to mark 0ish rows\columns
	var headRowZeros, headColZeros bool
	for i := range matrix[0] {
		if matrix[0][i] == 0 {
			headRowZeros = true
		}
	}
	for i := range matrix {
		if matrix[i][0] == 0 {
			headColZeros = true
		}
	}
	// check matrix and mark rows/colums to be filled with 0s in head row/col
	for r := len(matrix) - 1; r > 0; r-- {
		for c := len(matrix[r]) - 1; c > 0; c-- {
			if matrix[r][c] == 0 {
				matrix[r][0] = 0
				matrix[0][c] = 0
			}
		}
	}
	// actually fill rows and cols with 0s
	for r := len(matrix) - 1; r > 0; r-- {
		for c := len(matrix[r]) - 1; c > 0; c-- {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}
	if headRowZeros {
		for i := range matrix[0] {
			matrix[0][i] = 0
		}
	}
	if headColZeros {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}

}
