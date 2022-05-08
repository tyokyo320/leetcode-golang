/*
 * @lc app=leetcode id=220 lang=golang
 *
 * [220] Contains Duplicate III
 */

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	length := len(nums)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length && j-i <= k; j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}

	return false
}

func abs(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}

// @lc code=end

