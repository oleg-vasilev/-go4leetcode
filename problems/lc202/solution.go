package lc202

// 202. Happy Number

// Write an algorithm to determine if a number n is happy.
//
// A happy number is a number defined by the following process:
//
// Starting with any positive integer, replace the number
// by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay),
// or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.
//

func isHappy(n int) bool {
	var known = make(map[int]bool)
	for {
		next := 0
		// computing next value
		for n > 0 {
			next += (n % 10) * (n % 10)
			n = n / 10
		}
		// compare if we already know where it goes...
		if _, ok := known[next]; ok {
			// if we know - check if we happy
			return next == 1
		} else {
			known[next] = true
		}
		n = next
	}
}
