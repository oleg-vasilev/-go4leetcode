package lc121

import (
	"math"
	"slices"
)

// 121. Best Time to Buy and Sell Stock

// You are given an array prices where prices[i] is the price
// of a given stock on the ith day.
//
// You want to maximize your profit by choosing a single day to buy one
// stock and choosing a different day in the future to sell that stock.
//
// Return the maximum profit you can achieve from this transaction.
// If you cannot achieve any profit, return 0.

// memory and time inefficient approach
func maxProfit(prices []int) int {
	var bestPairs = make([]int, len(prices))

	if len(prices) < 2 {
		return 0
	}

	i1 := 0
	i2 := 1
	for i2 < len(prices) {
		bestPairs[i1] = max(bestPairs[i1], prices[i2]-prices[i1])
		if i2 < len(prices)-1 {
			i2++
		} else {
			i1++
			i2 = i1 + 1
		}
	}

	return slices.Max(bestPairs)
}

// time inefficient approach
func maxProfit1(prices []int) int {
	var profit int = 0

	if len(prices) < 2 {
		return 0
	}

	i1 := 0
	i2 := 1
	for i2 < len(prices) {
		profit = max(profit, prices[i2]-prices[i1])
		if i2 < len(prices)-1 {
			i2++
		} else {
			i1++
			i2 = i1 + 1
		}
	}
	return profit
}

// time inefficient approach
func maxProfit2(prices []int) int {
	var profit int = 0

	if len(prices) < 2 {
		return 0
	}

	i1 := 0
	i2 := 1
	for i2 < len(prices) {
		profit = max(profit, prices[i2]-prices[i1])
		if i2 < len(prices)-1 {
			i2++
		} else {
			i1++
			i2 = i1 + 1
		}
	}
	return profit
}

// basic dp
func maxProfit3(prices []int) int {
	lowestPrice := math.MaxInt
	bestProfit := 0

	var oldLowestPrice, oldBestProfit int
	for _, price := range prices {
		oldLowestPrice = lowestPrice
		oldBestProfit = bestProfit

		// update best profit
		bestProfit = max(oldBestProfit, price-oldLowestPrice)

		// update lowest price
		lowestPrice = min(oldLowestPrice, price)

	}
	return bestProfit
}

// optimized dp
func maxProfit4(prices []int) int {
	lowestPrice := math.MaxInt
	bestProfit := 0

	for _, price := range prices {
		bestProfit = max(bestProfit, price-lowestPrice)
		lowestPrice = min(lowestPrice, price)
	}
	return bestProfit
}

// two pointers
func maxProfitPointers(prices []int) int {
	bestProfit := 0
	for left, right := 0, 1; right < len(prices); right++ {
		if prices[left] > prices[right] {
			left = right
		} else {
			bestProfit = max(bestProfit, prices[right]-prices[left])
		}
	}
	return bestProfit
}
