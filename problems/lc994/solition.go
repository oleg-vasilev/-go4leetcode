// Package lc994 - Rotting oranges
package lc994

const (
	empty  = 0
	fresh  = 1
	rotten = 2
)

func orangesRotting(grid [][]int) int {
	currentState := grid
	nextState := make([][]int, len(grid))
	freshLeft := 0
	// fill buffer array with rows and also count fresh oranges
	for row := range grid {
		nextState[row] = make([]int, len(grid[row]))
		for col := range grid[row] {
			if grid[row][col] == fresh {
				freshLeft++
			}
		}
	}
	var stateChanged bool
	var rounds int
	for { // iterate round by round
		rounds++
		for row := 0; row < len(currentState); row++ {
			for col := 0; col < len(currentState[row]); col++ {
				switch currentState[row][col] {
				case empty:
					nextState[row][col] = empty
				case rotten:
					nextState[row][col] = rotten
				case fresh:
					// check neighbours - if at least one neighbour is rotten..
					if (row > 0 && currentState[row-1][col] == rotten) || // top
						(col > 0 && currentState[row][col-1] == rotten) || // left
						(row < len(currentState)-1 && currentState[row+1][col] == rotten) || // bottom
						(col < len(currentState[row])-1 && currentState[row][col+1] == rotten) { // right
						// orange becomes rotten too
						nextState[row][col] = rotten
						stateChanged = true
						freshLeft--
					} else {
						nextState[row][col] = fresh
					}
				}
			}
		}
		// update current state and use old field buffer as new state for next iteration
		currentState, nextState = nextState, currentState
		if !stateChanged { // no new rottens - stop
			break
		}
		stateChanged = false
	}
	if freshLeft == 0 {
		return rounds - 1
	} else {
		return -1
	}
}
