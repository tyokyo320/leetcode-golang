/*
 * @lc app=leetcode id=953 lang=golang
 *
 * [953] Verifying an Alien Dictionary
 */

// @lc code=start
func isAlienSorted(words []string, order string) bool {
	dict := make(map[int32]int)
	for k, v := range order {
		dict[v] = k
	}
	for i := 1; i < len(words); i++ {
		if !alienBigger(words[i], words[i-1], dict) {
			return false
		}
	}
	return true
}

func alienBigger(a, b string, dict map[int32]int) bool {
	la, lb := len(a), len(b)
	l := int(math.Min(float64(la), float64(lb)))
	for i := 0; i < l; i++ {
		if dict[int32(a[i])] > dict[int32(b[i])] {
			return true
		} else if dict[int32(a[i])] < dict[int32(b[i])] {
			return false
		}
		if len(a) == len(b) && dict[int32(a[i])] == dict[int32(b[i])] {
			return true
		}
	}
	if la > lb {
		return true
	}
	return false
}

// @lc code=end

