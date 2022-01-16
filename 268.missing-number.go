/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {

	length := len(nums)
	dict := make(map[int]bool)
	var result int

	for _, v := range nums {
		dict[v] = true
	}

	for i := 0; i <= length; i++ {
		if dict[i] == true {
			continue
		} else {
			result = i
		}
	}

	return result
}

// @lc code=end