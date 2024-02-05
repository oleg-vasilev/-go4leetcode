package lc120

/**

#120 - Triangle

Given a triangle array, return the minimum path sum from top to bottom.
For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the current row, you may move to either index i or index i + 1 on the next row.

Example 1:
Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
Output: 11
Explanation: The triangle looks like:
   2
  3 4
 6 5 7
4 1 8 3
The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).

Example 2:
Input: triangle = [[-10]]
Output: -10

*/

func minimumTotal(triangle [][]int) int {
	var minPath int
	for row := 0; row < len(triangle); row++ {
		for col := 0; col < len(triangle[row]); col++ {
			// check both possible ancestors of current cell and add smallest of them to current cell
			if row == 0 {
				minPath = triangle[row][col]
				// root item has no predecessors
				continue
			}
			// due to triangle form
			if col == 0 {
				// left corner case - only one ancestor
				right := triangle[row-1][col]
				triangle[row][col] += right
			} else if col == len(triangle[row])-1 {
				// right corner case - only one ancestor
				left := triangle[row-1][col-1]
				triangle[row][col] += left
			} else {
				// middle point - two ancestors
				left := triangle[row-1][col-1]
				right := triangle[row-1][col]
				if left < right {
					triangle[row][col] += left
				} else {
					triangle[row][col] += right
				}
			}
			// additionally track min path on last row
			if row == len(triangle)-1 {
				if col == 0 {
					minPath = triangle[row][col]
				} else {
					if minPath > triangle[row][col] {
						minPath = triangle[row][col]
					}
				}
			}
		}
	}
	return minPath
}
