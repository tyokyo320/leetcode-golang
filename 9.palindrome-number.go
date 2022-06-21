/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {

	// 当 x < 0 或者除了 0 以外的数字最后一位为 0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 偶数 || 奇数
	return x == revertedNumber || x == revertedNumber/10

}

// @lc code=end

