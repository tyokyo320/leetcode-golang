/*
 * @lc app=leetcode id=1768 lang=golang
 *
 * [1768] Merge Strings Alternately
 */

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {
	length1 := len(word1)
	length2 := len(word2)
	length := 0
	pointer := 0
	mergeStrings := ""

	if length1 < length2 {
		length = length1
	} else {
		length = length2
	}

	for i := 0; i < length; i++ {
		mergeStrings += string(word1[i]) + string(word2[i])
		pointer = i + 1
	}

	for pointer < length1 {
		mergeStrings += string(word1[pointer])
		pointer += 1
	}
	for pointer < length2 {
		mergeStrings += string(word2[pointer])
		pointer += 1
	}

	return mergeStrings

	// length1 := len(word1)
	// length2 := len(word2)
	// length := 0
	// result := make([]byte, length1+length2, length1+length2)

	// if length1 < length2 {
	// 	length = length1
	// 	for i, v := range word2[length:] {
	// 		result[length*2+i+1] = byte(v)
	// 	}
	// } else {
	// 	length = length2
	// 	for i, v := range word2[length:] {
	// 		result[length*2+i+1] = byte(v)
	// 	}
	// }

	// for i := 0; i < length; i++ {
	// 	result[i] = word1[i]
	// 	result[i+1] = word2[i+1]
	// }

	// return string(result[:])
}

// @lc code=end

