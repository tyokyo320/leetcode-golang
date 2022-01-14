/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {

	result := []int{}
	length := len(numbers)
	i := 0
	j := 1

	for i < length-1 {
		if numbers[i]+numbers[j] == target {
			result = append(result, i+1, j+1)
			return result
		}

		if (numbers[i]+numbers[j] > target) || (j+1 == length) {
			// back
			i += 1
			j = i
		}

		j += 1
	}
	return result
}

// @lc code=end

