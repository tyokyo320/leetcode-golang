/*
 * @lc app=leetcode id=404 lang=golang
 *
 * [404] Sum of Left Leaves
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
func sumOfLeftLeaves(root *TreeNode) int {
	return dfs(root)
}

func dfs(node *TreeNode) (result int) {

	if node == nil {
		return 0
	}

	if node.Left != nil {
		if isLeafNode(node.Left) {
			// 是叶子节点并且是某个节点的左子节点
			result += node.Left.Val
		} else {
			// 不是叶子节点要继续向左找
			result += dfs(node.Left)
		}
	}

	if node.Right != nil && !isLeafNode(node.Right) {
		result += dfs(node.Right)
	}

	return result
}

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

// @lc code=end

