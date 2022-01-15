/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	// 边界处理 例如 head = [1, 2], n = 2(n = 1)
	// 可以在head前一个结点，加一个newHead，slow，fast从newHead出发
	newHead := &ListNode{
		Next: head,
	}
	slow := newHead
	fast := newHead

	// fast先移动n step
	for n > 0 {
		fast = fast.Next
		n -= 1
	}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return newHead.Next
}

// @lc code=end

