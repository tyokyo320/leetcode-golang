/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	row, col := len(matrix), len(matrix[0])
	i, j := row-1, 0

	for i >= 0 && j < col {
		if target == matrix[i][j] {
			return true
		}
		if target < matrix[i][j] {
			i--
		} else {
			j++
		}
	}

	return false
}

// @lc code=end

