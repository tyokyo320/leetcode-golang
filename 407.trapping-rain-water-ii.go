/*
 * @lc app=leetcode id=407 lang=golang
 *
 * [407] Trapping Rain Water II
 */

// @lc code=start
func trapRainWater(heightMap [][]int) int {

	var volume int
	m := len(heightMap)
	n := len(heightMap[0])

	volume = bfs(heightMap, m, n)

	return volume
}

type block struct {
	x       int
	y       int
	height  int
	visited bool
}

func bfs(heightMap [][]int, m int, n int) int {
	var queue []*block
	waterHeight := 0
	sum := 0

	// 将四周先放入queue
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == m-1 || j == n-1 {
				block := &block{
					x:       i,
					y:       j,
					height:  heightMap[i][j],
					visited: true,
				}
				queue = append(queue, block)
			}
		}
	}

	for len(queue) > 0 {
		// update height
		current := queue[0]
		queue = queue[1:]
		waterHeight = max(waterHeight, current.height)
		// fmt.Printf("(%d, %d), height = %d, visited = %v\n", current.x, current.y, current.height, current.visited)

		for i := 0; i < 4; i++ {
			direction := []int{1, 0, -1, 0, 1}
			x := direction[i]
			y := direction[i+1]

			if current.x+x >= 0 && current.x+x < m && current.y+y >= 0 && current.y+y < n && current.visited == false {
				fmt.Printf("(%d, %d), height = %d, visited = %v\n", current.x, current.y, current.height, current.visited)
				if current.height < waterHeight {
					sum += (waterHeight - current.height)
				}

				b := &block{
					x:       current.x + x,
					y:       current.y + y,
					height:  heightMap[current.x+x][current.y+y],
					visited: true,
				}
				queue = append(queue, b)
			}
		}
	}

	return sum

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

