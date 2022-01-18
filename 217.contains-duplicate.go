/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {

	dict := make(map[int]bool)

	for _, v := range nums {
		// if not include num, insert into map
		if !dict[v] {
			dict[v] = true
		} else {
			return true
		}
	}
	return false
}

// @lc code=end

