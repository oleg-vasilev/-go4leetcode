package lc289

// 289. Game of Life

// According to Wikipedia's article: "The Game of Life,
// also known simply as Life, is a cellular automaton
// devised by the British mathematician John Horton Conway in 1970."
//
// The board is made up of an m x n grid of cells, where each cell
// has an initial state: live (represented by a 1) or dead (represented by a 0).
// Each cell interacts with its eight neighbors
// (horizontal, vertical, diagonal) using the following
// four rules (taken from the above Wikipedia article):
//
// Any live cell with fewer than two live neighbors dies as if caused by under-population.
// Any live cell with two or three live neighbors lives on to the next generation.
// Any live cell with more than three live neighbors dies, as if by over-population.
// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

// The next state is created by applying the above rules simultaneously to
// every cell in the current state, where births and deaths occur
// simultaneously. Given the current state of the m x n grid board,
// return the next state.

func gameOfLife(board [][]int) {
	// Straightforward solution  with additional board allocation
	var nextState = make([][]int, len(board))
	for n := range nextState {
		nextState[n] = make([]int, len(board[n]))
	}

	var aliveNeighbours int
	for r := range board {
		for c := range board[r] {
			aliveNeighbours = 0
			for _, shift := range neighbours {
				if shift[0]+r < 0 || shift[0]+r >= len(board) {
					// row out of bound
					continue
				}
				if shift[1]+c < 0 || shift[1]+c >= len(board[r]) {
					// column out of bound
					continue
				}
				aliveNeighbours += board[shift[0]+r][shift[1]+c] // because live cells contains 1s
			}
			if board[r][c] == 1 && (aliveNeighbours < 2 || aliveNeighbours > 3) {
				nextState[r][c] = 0
				continue
			}
			if board[r][c] == 1 && (aliveNeighbours == 2 || aliveNeighbours == 3) {
				nextState[r][c] = 1
				continue
			}
			if board[r][c] == 0 && aliveNeighbours == 3 {
				nextState[r][c] = 1
				continue
			}
		}
	}

	for r := range board {
		copy(board[r], nextState[r])
	}

}

var neighbours = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1}}
