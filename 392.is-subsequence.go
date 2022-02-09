/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {

	// 双指针

	bs := []byte(s)
	bt := []byte(t)

	length1 := len(s)
	length2 := len(t)

	if length1 == 0 && length2 == 0 {
		return true
	}

	if length2 == 0 {
		return false
	}

	// pointer
	i := 0
	j := 0
	for i < length1 {
		if bs[i] == bt[j] {
			i += 1
		}
		j += 1
		if i < length1 && j == length2 {
			return false
		}
	}

	return true

	/*
		优化

		n, m := len(s), len(t)
		i, j := 0, 0
		for i < n && j < m {
			if s[i] == t[j] {
				i++
			}
			j++
		}
		return i == n
	*/

}

// @lc code=end

