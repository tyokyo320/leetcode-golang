/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {

	b1 := []byte(s)
	b2 := []byte(t)

	dict1 := make(map[byte]int)
	dict2 := make(map[byte]int)

	for _, v := range b1 {
		dict1[v] += 1
	}
	for _, v := range b2 {
		dict2[v] += 1
	}

	// 判断两个map是否相同
	if reflect.DeepEqual(dict1, dict2) {
		return true
	} else {
		return false
	}
}

// @lc code=end

