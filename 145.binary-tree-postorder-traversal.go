/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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
func postorderTraversal(root *TreeNode) []int {

	result := []int{}
	result = postorder(root, result)

	return result
}

func postorder(node *TreeNode, result []int) []int {

	if node == nil {
		return result
	}

	result = postorder(node.Left, result)
	result = postorder(node.Right, result)
	result = append(result, node.Val)

	return result
}

// @lc code=end

