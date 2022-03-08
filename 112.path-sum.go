/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	var currSum int = 0

	return dfs(root, currSum, targetSum)
}

func dfs(node *TreeNode, currSum, targetSum int) bool {
	if node == nil {
		return false
	}

	currSum += node.Val
	if currSum == targetSum && node.Left == nil && node.Right == nil {
		return true
	}

	return dfs(node.Left, currSum, targetSum) || dfs(node.Right, currSum, targetSum)
}

// @lc code=end

