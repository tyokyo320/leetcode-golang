/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	var result []int

	for i, v := range nums {
		if _, ok := dict[target-v]; ok {
			result = append(result, dict[target-v], i)
		}
		dict[v] = i
	}
	return result
}

// @lc code=end

