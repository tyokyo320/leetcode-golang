/*
 * @lc app=leetcode id=1502 lang=golang
 *
 * [1502] Can Make Arithmetic Progression From Sequence
 */

// @lc code=start
func canMakeArithmeticProgression(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	length := len(arr)

	d := arr[1] - arr[0]
	for i := 1; i < length; i++ {
		if arr[i]-arr[i-1] == d {
			continue
		} else {
			return false
		}
	}

	return true
}

// @lc code=end

