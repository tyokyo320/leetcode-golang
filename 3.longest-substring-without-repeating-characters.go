/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
import "fmt"

func lengthOfLongestSubstring(s string) int {

	b := []byte(s)
	length := len(b)
	sum := 1
	current := 1
	dict := make(map[byte]int)
	start := 0
	end := 0

	for end < length {

		// 无重复：更新map，移动end
		if dict[b[end]] != 1 {
			dict[b[end]] = 1
			current = end - start + 1
			sum = max(current, sum)
			// fmt.Printf("%d - %d = %d, sum = %d\n", end, start, current, sum)
			end += 1
			continue
		}

		// 有重复：先更新map，start移动至无重复位置
		for b[start] != b[end] {
			dict[b[start]] = 0
			start += 1
		}

		dict[b[start]] = 0
		start += 1
		// fmt.Printf("%d, %d\n", start, end)
	}

	if length == 0 {
		return 0
	} else {
		return sum
	}

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

