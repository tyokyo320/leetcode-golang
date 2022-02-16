/*
 * @lc app=leetcode id=566 lang=golang
 *
 * [566] Reshape the Matrix
 */

// @lc code=start
func matrixReshape(mat [][]int, r int, c int) [][]int {
	m := len(mat)
	n := len(mat[0])
	result := make([][]int, r)
	for i, _ := range result {
		result[i] = make([]int, c)
	}

	if m*n != r*c {
		return mat
	}

	var row int = 0
	var col int = 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[row][col] = mat[i][j]
			col += 1

			if col == c {
				row += 1
				col = 0
			}
		}
	}

	return result
}

// @lc code=end

