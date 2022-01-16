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

	start := 0
	end := 0
	initMap(dict, b1)

	for dict[b2[end]] != 1 {
		// start，end移动到第一个s1存在的字符位置
		end += 1
		start += 1
		continue
	}

	for end < length2 {

		dict[b2[end]] -= 1
		end += 1

		if isFinished(dict) {
			return true
		}
		if end >= length2 {
			return false
		}

		fmt.Printf("start = %d, end = %d\n", start, end)
		fmt.Println(dict)
		// fmt.Printf("start: %s, end: %s\n", string(b2[start]), string(b2[end]))

		// 不含s1字符
		if _, ok := dict[b2[end]]; !ok {
			end += 1
			start = end
			initMap(dict, b1)
			continue
		}

		// 含s1字符且重复
		if _, ok := dict[b2[end]]; ok {
			if isDuplicate(b2[start], b2[end]) {
				dict[b2[start]] += 1
				start += 1
			}
		}

		// 含s1字符且不重复
		if _, ok := dict[b2[end]]; ok {
			if !isDuplicate(b2[start], b2[end]) {
				if (end - start + 1) == length1 {
					if isFinished(dict) {
						return true
					}
				} else {
					fmt.Printf("--- start = %d, end = %d ---\n", start, end)
					fmt.Println(string(b2[end]))
					dict[b2[end]] -= 1
					fmt.Println(dict)
				}
			}
			continue
		}
	}
	return isFinished(dict)
}

// initMap 将s1字符保存至dict
func initMap(dict map[byte]int, b1 []byte) {
	for _, v := range b1 {
		dict[v] += 1
	}
}

func isDuplicate(a byte, b byte) bool {
	if a == b {
		return true
	} else {
		return false
	}
}

func isFinished(dict map[byte]int) bool {
	for _, v := range dict {
		if v <= 0 {
			continue
		} else {
			return false
		}
	}

	return true
}

// @lc code=end

