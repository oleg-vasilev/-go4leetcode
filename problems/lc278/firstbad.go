package lc278

// #278 - First Bad Version

// You are a product manager and currently leading a team to develop a new product.
// Unfortunately, the latest version of your product fails the quality check.
// Since each version is developed based on the previous version,
// all the versions after a bad version are also bad.
//
// Suppose you have n versions [1, 2, ..., n] and you want to find out the
// first bad one, which causes all the following ones to be bad.
//
// You are given an API bool isBadVersion(version) which returns whether
// version is bad. Implement a function to find the first bad version.
// You should minimize the number of calls to the API.

type versionBlackBox struct {
	badVersion int
}

func (bb versionBlackBox) isBadVersion(version int) bool {
	return version >= bb.badVersion
}

func firstBadVersion(n int, bb versionBlackBox) int {
	var minBad = n
	var maxGood, curr int
	for {
		curr = (minBad + maxGood) / 2
		if minBad-maxGood == 1 {
			return minBad
		}
		if bb.isBadVersion(curr) && curr < minBad {
			minBad = curr
		} else if curr > maxGood {
			maxGood = curr
		}
	}
}
