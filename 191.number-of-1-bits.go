/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
func hammingWeight(num uint32) int {

	/*
		对于只有一个bit位为1的值来说其与其本身减1的值与运算后必为0；
		该思路一般化后对于任意的数，只要不断将其与其本身减1后的值与运算直至为0即可得到bit位为1的数量；

		与运算符（&）：两个同时为1，结果为1，否则为0
		或运算（|）：参加运算的两个对象，一个为1，其值为1
		异或运算符（^）：参加运算的两个对象，如果两个位为“异”（值不同），则该位结果为1，否则为0
	*/
	var count int = 0

	for num != 0 {
		num &= num - 1
		count += 1
	}

	return count
}

// @lc code=end

