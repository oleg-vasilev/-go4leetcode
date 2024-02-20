package lc122

// 122. Best Time to Buy and Sell Stock II

// You are given an integer array prices where prices[i] is the price
// of a given stock on the ith day.
//
// On each day, you may decide to buy and/or sell the stock.
// You can only hold at most one share of the stock at any time.
// However, you can buy it then immediately sell it on the same day.
//
// Find and return the maximum profit you can achieve.

// just math solution using fact that we can sell and buy at same day,
// so we can just stonks every time price will rises

func maxProfit(prices []int) int {
	var profit int

	for i, _ := range prices {
		if i == 0 {
			continue
		}
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1] // stonks every time we know price will rise
		}
	}
	return profit
}
