/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
import "fmt"

func checkInclusion(s1 string, s2 string) bool {

	b1 := []byte(s1)
	b2 := []byte(s2)
	length1 := len(b1)
	length2 := len(b2)
	dict := make(map[byte]int)
	slidingWindow := make(map[byte]int)

	initMap(dict, b1)
	// s1字符串字符的种类数
	sCount := len(dict)
	// sliding-window中字符出现的种类数，当滑动窗口中某个字符个数与s1中对应字符相等时才计数
	swCount := 0

	left := 0
	right := 0

	// 先让sliding-window的右边界向右移动
	for right < length2 {
		// s1中存在的字符才统计
		if dict[b2[right]] > 0 {
			// 滑动窗口中出现的字符的频数更新
			slidingWindow[b2[right]] += 1
			// 滑动窗口中出现的右边界的次数与s1当中对应的字符频数相等时，计数
			if slidingWindow[b2[right]] == dict[b2[right]] {
				swCount += 1
			}
		}
		// 右边界对应读取了字符后移动
		right += 1
		// fmt.Println(sCount, swCount)

		// 左边界移动条件：当滑动窗口出现字符的种类数 == s1中字符种类数，并且对应频数相等
		for sCount == swCount {
			// fmt.Println(left, right)
			if right-left == length1 {
				return true
			}
			// s1中存在的字符才统计
			if dict[b2[left]] > 0 {
				// 滑动窗口中出现的字符的频数更新
				slidingWindow[b2[left]] -= 1
				// 滑动窗口中出现的左边界的次数小于s1当中对应的字符频数时
				if slidingWindow[b2[left]] < dict[b2[left]] {
					swCount -= 1
				}
			}
			left += 1
		}
	}
	return false
}

// initMap 统计s1字符的频数
func initMap(dict map[byte]int, b1 []byte) {
	for _, v := range b1 {
		dict[v] += 1
	}
}

// @lc code=end

