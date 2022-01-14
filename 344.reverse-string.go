/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */

// @lc code=start
func reverseString(s []byte) {
	length := len(s)
	reverse(s, 0, length-1)
}

func reverse(s []byte, i int, j int) {
	for i <= j {
		s[i], s[j] = s[j], s[i]
		i += 1
		j -= 1
	}
}

// @lc code=end

