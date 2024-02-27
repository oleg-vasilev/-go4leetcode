package lc56

import (
	"sort"
)

// 56. Merge Intervals

// Given an array of intervals where intervals[i] = [starti, endi],
// merge all overlapping intervals, and return an array of the
// non-overlapping intervals that cover all the intervals in the input.

// Straightforward logic without sorting
func merge(intervals [][]int) [][]int {
	var result = make([][]int, 0, len(intervals))

	var merged = make(map[int]bool)
	for a := 0; a < len(intervals); a++ {
		// skip merged values
		if merged[a] {
			continue
		}
		intA := Interval(intervals[a])
		// looking for interval B that A can consume
		for b := a + 1; b < len(intervals); b++ {
			// skip merged values
			if merged[b] {
				continue
			}
			intB := Interval(intervals[b])
			// B is merged to A and should be skipped in future
			if intA.Merge(intB) {
				merged[b] = true
				// if b is merged the a should make another loop from beginning
				// because now A can possibly consume something else
				b = a
				continue
			}
		}
		// now A includes all possible intervals and can be added to result
		result = append(result, intA[:])
	}
	return result
}

type Interval [2]int

func (in *Interval) start() int {
	return in[0]
}

func (in *Interval) end() int {
	return in[1]
}

func (in *Interval) Merge(other Interval) bool {
	// two non overlapping cases
	if in.start() > other.end() || in.end() < other.start() {
		return false
	}
	// overlapping case in any overlapping
	in[0] = min(in.start(), other.start())
	in[1] = max(in.end(), other.end())
	return true
}

// much more effective approach with sorting intervals by start value
func mergeWithSort(intervals [][]int) [][]int {
	var results [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, interval := range intervals {
		lastResult := len(results) - 1
		if len(results) != 0 && results[lastResult][1] >= interval[0] {
			// last result consumes interval
			results[lastResult][1] = max(results[lastResult][1], interval[1])
		} else {
			// new interval in result
			results = append(results, interval)
		}
	}
	return results
}
