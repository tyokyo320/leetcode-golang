/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}

	// BFS
	for i := 0; len(queue) > 0; i++ {
		result = append(result, []int{})

		for j := len(queue); j > 0; j-- {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}

			result[i] = append(result[i], queue[0].Val)
			// deque
			queue = queue[1:]
		}
	}

	return result
}

// @lc code=end

