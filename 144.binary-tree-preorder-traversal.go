/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
func preorderTraversal(root *TreeNode) []int {

	result := []int{}
	result = preorder(root, result)

	return result
}

func preorder(node *TreeNode, result []int) []int {

	if node == nil {
		return result
	}

	result = append(result, node.Val)
	result = preorder(node.Left, result)
	result = preorder(node.Right, result)

	return result
}

// @lc code=end

