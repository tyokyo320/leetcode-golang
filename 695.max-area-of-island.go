/*
 * @lc app=leetcode id=695 lang=golang
 *
 * [695] Max Area of Island
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	maxArea := 0
	// 初始化2维数组
	cache := make([][]int, m)
	for i, _ := range cache {
		cache[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxArea = max(maxArea, dfs(grid, i, j, m, n, cache))
		}
	}
	return maxArea
}

func dfs(grid [][]int, i int, j int, m int, n int, cache [][]int) int {

	// out of bounds or no island
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
		return 0
	}

	if cache[i][j] != 0 {
		return cache[i][j]
	}

	// current
	area := 1
	grid[i][j] = 0

	// left, right, top, down
	area += dfs(grid, i, j-1, m, n, cache) + dfs(grid, i, j+1, m, n, cache) + dfs(grid, i-1, j, m, n, cache) + dfs(grid, i+1, j, m, n, cache)
	cache[i][j] = area
	return area
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end