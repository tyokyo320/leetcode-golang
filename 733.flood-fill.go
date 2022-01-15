/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	m := len(image)
	n := len(image[0])

	bfs(image, sr, sc, newColor, m, n)

	return image
}

type point struct {
	x int
	y int
}

func bfs(image [][]int, sr int, sc int, newColor int, m int, n int) {

	var queue []*point

	color := image[sr][sc]

	if color == newColor {
		return
	}

	initPoint := &point{
		x: sr,
		y: sc,
	}
	queue = append(queue, initPoint)

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		image[temp.x][temp.y] = newColor

		// check 4-directionally
		// 上: (sr-1, sc)
		// 下: (sr+1, sc)
		// 左: (sr, sc-1)
		// 右: (sr, sc+1)
		direction := []int{1, 0, -1, 0, 1}
		for i := 0; i < 4; i++ {
			x := direction[i]
			y := direction[i+1]

			if (temp.x+x >= 0) && (temp.x+x < m) && (temp.y+y >= 0) && (temp.y+y < n) && image[temp.x+x][temp.y+y] == color {
				point := &point{
					x: temp.x + x,
					y: temp.y + y,
				}
				queue = append(queue, point)
			}
		}
	}

}

// @lc code=end

