/*
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
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
func findTarget(root *TreeNode, k int) bool {
	dict := make(map[int]bool)
	return find(root, k, dict)
}

func find(node *TreeNode, k int, dict map[int]bool) bool {
	if node == nil {
		return false
	}

	if _, ok := dict[k-node.Val]; ok {
		return true
	}

	dict[node.Val] = true
	return find(node.Left, k, dict) || find(node.Right, k, dict)
}

// @lc code=end

