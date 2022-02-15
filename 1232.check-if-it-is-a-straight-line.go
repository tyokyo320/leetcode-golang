/*
 * @lc app=leetcode id=1232 lang=golang
 *
 * [1232] Check If It Is a Straight Line
 */

// @lc code=start
func checkStraightLine(coordinates [][]int) bool {
	length := len(coordinates)
	x0, x1 := coordinates[0][0], coordinates[1][0]
	y0, y1 := coordinates[0][1], coordinates[1][1]

	for i := 1; i < length; i++ {
		x2, y2 := coordinates[i][0], coordinates[i][1]
		if (x1-x0)*(y2-y1) != (y1-y0)*(x2-x1) {
			return false
		}
	}
	return true
}

// @lc code=end

