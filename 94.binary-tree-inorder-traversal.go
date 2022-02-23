/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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
func inorderTraversal(root *TreeNode) []int {

	result := []int{}
	result = inorder(root, result)

	return result
}

func inorder(node *TreeNode, result []int) []int {

	if node == nil {
		return result
	}

	result = inorder(node.Left, result)
	result = append(result, node.Val)
	result = inorder(node.Right, result)

	return result
}

// @lc code=end

