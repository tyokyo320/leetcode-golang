/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	bRansomNote := []byte(ransomNote)
	bMagazine := []byte(magazine)
	lengthNote := len(bRansomNote)
	index := 0

	dict := make(map[byte]int)

	for _, v := range bMagazine {
		dict[v] += 1
	}

	for index < lengthNote {
		// 若note中存在magazine的字符
		if _, ok := dict[bRansomNote[index]]; ok {
			dict[bRansomNote[index]] -= 1
			// node中对应字符的数量多于magazine
			if dict[bRansomNote[index]] < 0 {
				return false
			}
		} else {
			// 不存在magazine的字符，说明无法构成
			return false
		}
		index += 1
	}

	return true
}

// @lc code=end

