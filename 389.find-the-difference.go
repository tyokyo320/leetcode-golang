/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	length := len(t)
	sum1 := 0
	sum2 := 0

	for i := 0; i < length; i++ {
		sum2 += int(t[i])
	}
	for i := 0; i < length-1; i++ {
		sum1 += int(s[i])
	}

	return byte(sum2 - sum1)
}

// @lc code=end

