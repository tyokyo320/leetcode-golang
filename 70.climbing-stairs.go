/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 2

	// 第n阶只能有n-1阶跨1阶或者n-2阶跨2阶达到：f(n) = f(n - 1) + f(n - 2)
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	// fmt.Println(dp)

	return dp[n-1]
}

// @lc code=end