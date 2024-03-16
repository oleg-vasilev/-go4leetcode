package lc57

// 57. Insert Interval

// You are given an array of non-overlapping intervals intervals where
// intervals[i] = [starti, endi] represent the start and the end of the
// ith interval and intervals is sorted in ascending order by starti.
// You are also given an interval newInterval = [start, end] that
// represents the start and end of another interval.
//
// Insert newInterval into intervals such that intervals is still sorted
// in ascending order by starti and intervals still does not have any
// overlapping intervals (merge overlapping intervals if necessary).
//
// Return intervals after the insertion.
//
// Note that you don't need to modify intervals in-place.
// You can make a new array and return it.
//

func insert(intervals [][]int, newInterval []int) [][]int {
	var res = make([][]int, 0, len(intervals)+1)

	var inserted bool
	for _, curr := range intervals {
		// non overlapping case #1 current interval is definitely before new
		if curr[1] < newInterval[0] {
			res = append(res, curr)
			continue
		}
		// non overlapping case #2 current interval is definitely after new
		if curr[0] > newInterval[1] {
			// if new isn't inserted yet - insert it
			if !inserted {
				res = append(res, newInterval)
				inserted = true
			}
			// and add current afterwards
			res = append(res, curr)
			continue
		}
		// overlapping cases - megre current onto new to handle new bounds in next iterations
		newInterval[0] = min(newInterval[0], curr[0])
		newInterval[1] = max(newInterval[1], curr[1])
	}
	// handle corner cases
	if len(res) == 0 || !inserted {
		res = append(res, newInterval)
	}
	return res
}
