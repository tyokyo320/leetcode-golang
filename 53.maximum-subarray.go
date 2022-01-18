/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
import "fmt"

func maxSubArray(nums []int) int {

	length := len(nums)

	if length == 1 {
		return nums[0]
	}

	dp := make([]int, length)
	result := nums[0]

	dp[0] = nums[0]
	// fmt.Println(dp)

	for i := 1; i < length; i++ {
		// max(选，不选)
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(result, dp[i])
		// fmt.Printf("i = %d, dp[i] = %d\n", i, dp[i])
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

