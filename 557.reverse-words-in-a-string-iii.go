/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */

// @lc code=start
import "fmt"

func reverseWords(s string) string {
	space := 0
	i := 0
	j := 0
	bytes := []byte(s)
	// 方法1：加一个空白
	bytes = append(bytes, ' ')
	length := len(bytes)

	for j < length {
		if bytes[j] != ' ' {
			j += 1
			// 方法2：处理最后一个单词没换位
			// if j == length {
			// 	reverse(bytes[i:j], 0, j-1-i)
			// }
			continue
		}

		// space position
		space = j
		// 这里切了之后变成了新的list，传入i,j-1会越界
		// 只需要传入每个单词的start和end即可
		reverse(bytes[i:j], 0, j-1-i)
		i = space + 1
		j = i

	}
	// s = string(bytes)
	// 方法1：加完空白剪掉加的空白返回
	s = string(bytes[:length-1])
	return s
}

func reverse(bytes []byte, start int, end int) {
	for start <= end {
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start += 1
		end -= 1
	}
}

// @lc code=end

