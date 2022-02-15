/*
 * @lc app=leetcode id=1588 lang=golang
 *
 * [1588] Sum of All Odd Length Subarrays
 */

// @lc code=start
func sumOddLengthSubarrays(arr []int) int {
	length := len(arr)
	odds := makeOdd(length)
	result := 0

	for _, odd := range odds {
		for i := 0; i < length; i++ {
			if odd > length {
				break
			}
			result += sums(arr[i:odd])
			odd += 1
		}
	}

	return result
}

func sums(arr []int) int {
	var sums int = 0

	for i := 0; i < len(arr); i++ {
		sums += arr[i]
	}

	return sums
}

func makeOdd(n int) []int {
	odds := []int{}

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			odds = append(odds, i)
		}
	}

	return odds
}

// @lc code=end

