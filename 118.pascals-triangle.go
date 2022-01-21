/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
import "fmt"

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}

	for i := 1; i < numRows; i++ {
		// 第一个元素和最后一个元素始终为1
		dp := make([]int, i+1)
		dp[0], dp[i] = 1, 1

		// 中间元素利用上一排元素两辆作和得到
		for j := 1; j < i; j++ {
			dp[j] = result[i-1][j-1] + result[i-1][j]
		}
		result[i] = dp
	}
	return result
}

// @lc code=end

