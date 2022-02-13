/*
 * @lc app=leetcode id=1779 lang=golang
 *
 * [1779] Find Nearest Point That Has the Same X or Y Coordinate
 */

// @lc code=start
func nearestValidPoint(x int, y int, points [][]int) int {
	i := len(points)
	manhattanDistance := 10000
	index := -1
	temp := 0

	for j := 0; j < i; j++ {
		// valid point
		if x == points[j][0] || y == points[j][1] {
			temp = manhattanDistance
			// manhattanDistance = min(manhattanDistance, int(math.Abs(float64(x-points[j][0])))+int(math.Abs(float64(y-points[j][1]))))
			manhattanDistance = min(manhattanDistance, abs(x-points[j][0])+abs(y-points[j][1]))
			if manhattanDistance < temp {
				index = j
			}
		}
	}

	return index
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// @lc code=end

