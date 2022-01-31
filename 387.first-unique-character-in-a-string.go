/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func firstUniqChar(s string) int {

	b := []byte(s)
	dict := make(map[byte]int)

	for _, v := range b {
		dict[v] += 1
	}

	index := 0
	for _, value := range b {
		if dict[value] == 1 {
			return index
		}
		index += 1
	}

	return -1
}

// @lc code=end

