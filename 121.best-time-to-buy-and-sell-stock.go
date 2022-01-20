/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
import "fmt"

func maxProfit(prices []int) int {
	length := len(prices)
	// lowest profit up to i-th day
	minPrice := make([]int, length)
	// max profit up to i-th day
	maxProfit := make([]int, length)

	minPrice[0] = prices[0]
	maxProfit[0] = 0

	for i := 1; i < length; i++ {
		minPrice[i] = min(minPrice[i-1], prices[i])
		// max(不操作，当天卖出：当天价格-前i-1天最低价格)
		maxProfit[i] = max(maxProfit[i-1], prices[i]-minPrice[i])
	}

	return maxProfit[length-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

