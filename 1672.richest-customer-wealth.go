/*
 * @lc app=leetcode id=1672 lang=golang
 *
 * [1672] Richest Customer Wealth
 */

// @lc code=start
func maximumWealth(accounts [][]int) int {
	length := len(accounts)
	richestWealth := 0
	wealth := 0

	for i := 0; i < length; i++ {
		wealth = sums(accounts[i])
		richestWealth = max(richestWealth, wealth)
	}

	return richestWealth
}

func sums(arr []int) int {
	var sums int = 0

	for i := 0; i < len(arr); i++ {
		sums += arr[i]
	}

	return sums
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// @lc code=end

