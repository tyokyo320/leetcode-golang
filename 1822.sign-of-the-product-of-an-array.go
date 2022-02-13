/*
 * @lc app=leetcode id=1822 lang=golang
 *
 * [1822] Sign of the Product of an Array
 */

// @lc code=start
func arraySign(nums []int) int {
	sign := 1

	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			sign = -sign
		}
	}

	return sign
}

// @lc code=end

