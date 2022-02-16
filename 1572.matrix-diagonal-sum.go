/*
 * @lc app=leetcode id=1572 lang=golang
 *
 * [1572] Matrix Diagonal Sum
 */

// @lc code=start
func diagonalSum(mat [][]int) int {
	/*
		length := len(mat)
		var sums int = 0

		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				if (i == j) || (i+j == length-1) {
					sums += mat[i][j]
				}
			}
		}

		return sums
	*/
	length := len(mat)
	var sums int = 0
	mid := length / 2

	for i := 0; i < length; i++ {
		sums += mat[i][i] + mat[i][length-1-i]
	}

	// 奇数时，中间会被多计算一次
	if length%2 == 1 {
		return sums - mat[mid][mid]
	} else {
		return sums
	}
}

// @lc code=end

