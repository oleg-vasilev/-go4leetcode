package lc274

import (
	"slices"
)

// 274. H-Index

// Given an array of integers citations where citations[i] is the number
// of citations a researcher received for their ith paper,
// return the researcher's h-index.
//
// According to the definition of h-index on Wikipedia:
// The h-index is defined as the maximum value of h such that the
// given researcher has published at least h papers that have each
// been cited at least h times.
//

func hIndex(citations []int) int {
	slices.Sort(citations)
	slices.Reverse(citations)
	var h int
	for i, v := range citations {
		if i+1 <= v {
			h = i + 1
		}

	}
	return h
}
