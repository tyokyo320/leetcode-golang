/*
 * @lc app=leetcode id=1523 lang=golang
 *
 * [1523] Count Odd Numbers in an Interval Range
 */

// @lc code=start
func countOdds(low int, high int) int {

	if low%2 == 1 || high%2 == 1 {
		return (high-low)/2 + 1
	} else {
		return (high - low) / 2
	}
}

// @lc code=end

