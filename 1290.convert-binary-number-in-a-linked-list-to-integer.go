/*
 * @lc app=leetcode id=1290 lang=golang
 *
 * [1290] Convert Binary Number in a Linked List to Integer
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	result := 0

	for head != nil {
		result = result*2 + head.Val
		head = head.Next
	}

	return result
}

// @lc code=end
