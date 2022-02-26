/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	dict := make(map[string]bool, 9)

	return checkRow(board, dict) && checkCol(board, dict) && checkNineBlock(board, dict)
}

func initMap(dict map[string]bool) {
	for i := 1; i < 10; i++ {
		// int转字符串
		dict[strconv.Itoa(i)] = false
	}
}

func checkRow(board [][]byte, dict map[string]bool) bool {
	row, col := len(board), len(board[0])
	initMap(dict)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if string(board[i][j]) == "." {
				continue
			}

			// already in map
			if dict[string(board[i][j])] == true {
				return false
			}
			// insert into map
			dict[string(board[i][j])] = true
		}
		initMap(dict)
	}

	return true
}

func checkCol(board [][]byte, dict map[string]bool) bool {
	row, col := len(board), len(board[0])
	initMap(dict)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if string(board[j][i]) == "." {
				continue
			}

			// already in map
			if dict[string(board[j][i])] == true {
				return false
			}
			// insert into map
			dict[string(board[j][i])] = true
		}
		initMap(dict)
	}

	return true
}

// check 3 * 3
func checkNineBlock(board [][]byte, dict map[string]bool) bool {
	initMap(dict)

	// 将大的9宫格拆分为3*3的小九宫格
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			initMap(dict)
			for k := i * 3; k < i*3+3; k++ {
				for l := j * 3; l < j*3+3; l++ {
					if string(board[k][l]) == "." {
						continue
					}

					// already in map
					if dict[string(board[k][l])] == true {
						return false
					}
					// insert into map
					dict[string(board[k][l])] = true
				}
			}
		}
	}

	return true
}

// @lc code=end

