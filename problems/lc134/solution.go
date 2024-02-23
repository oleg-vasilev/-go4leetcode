package lc134

// 134. Gas Station

// There are n gas stations along a circular route, where the amount
// of gas at the ith station is gas[i].
//
// You have a car with an unlimited gas tank and it costs cost[i]
// of gas to travel from the ith station to its next (i + 1)th station.
// You begin the journey with an empty tank at one of the gas stations.
//
// Given two integer arrays gas and cost, return the starting gas
// station's index if you can travel around the circuit once in the
// clockwise direction, otherwise return -1. If there exists
// a solution, it is guaranteed to be unique

// Constraints:
//
// n == gas.length == cost.length
// 1 <= n <= 105
// 0 <= gas[i], cost[i] <= 104

// bruteforce simulation with bad time complexity
// func canCompleteCircuit(gas []int, cost []int) int {
// 	var tank int
// 	var startPos int
// 	var currPos int
// 	for {
// 		// fill the tank at current position
// 		tank += gas[currPos]
// 		// check if we can travel to next spot
// 		if tank >= cost[currPos] {
// 			// travel to next station
// 			tank -= cost[currPos]
// 			currPos++
// 			// loop position back to first
// 			if currPos == len(gas) {
// 				currPos = 0
// 			}
// 			if currPos == startPos {
// 				// we made a loop!
// 				return startPos
// 			}
// 		} else {
// 			// can't travel - select next starting pos
// 			if startPos < len(gas)-1 {
// 				startPos++
// 				currPos = startPos
// 				tank = 0
// 			} else {
// 				// we tried all starting positions
// 				return -1
// 			}
// 		}
// 	}
// }

// greedy with O(n) complexity
func canCompleteCircuit(gas []int, cost []int) int {

	var tank, startPos int
	var totalGas, totalDist int
	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalDist += cost[i]
		tank += gas[i] - cost[i]
		if tank < 0 {
			// if tank is negative - station is unreachable from any
			// other station between our start and current positions
			startPos = i + 1
			tank = 0
		}
	}
	// if total gas value is lesser then total loop distance
	// it is impossible to made a loop at all from any point
	if totalGas < totalDist {
		return -1
	} else {
		return startPos
	}
}
