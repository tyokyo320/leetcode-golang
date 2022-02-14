/*
 * @lc app=leetcode id=1790 lang=golang
 *
 * [1790] Check if One String Swap Can Make Strings Equal
 */

// @lc code=start
func areAlmostEqual(s1 string, s2 string) bool {
	length := len(s1)
	counter := 0
	pointer := 0
	dict1 := make(map[byte]bool)
	dict2 := make(map[byte]bool)

	for pointer < length {
		if s1[pointer] != s2[pointer] {
			// s1 = "qgqeg" s2 = "gqgeq"
			if _, ok := dict1[s1[pointer]]; ok {
				return false
			}
			// insert into map
			dict1[s1[pointer]] = true
			dict2[s2[pointer]] = true
		} else {
			// count the number of equal strings
			counter += 1
		}
		pointer += 1
		if len(dict1) > 2 || len(dict2) > 2 {
			return false
		}
		// two strings are already equal
		if counter == length {
			return true
		}
	}

	if len(dict1) == 2 && len(dict2) == 2 {
		return reflect.DeepEqual(dict1, dict2)
	} else {
		return false
	}
}

// @lc code=end

