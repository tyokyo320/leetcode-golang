/*
 * @lc app=leetcode id=1356 lang=golang
 *
 * [1356] Sort Integers by The Number of 1 Bits
 */

// @lc code=start
func sortByBits(arr []int) []int {

	sort.Slice(arr, func(i, j int) bool {
		x, y := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		if x == y {
			return arr[i] < arr[j]
		}
		return x < y
	})

	return arr
}

// // convertToBin 10进制转2进制字符串
// func convertToBin(num int) string {
// 	s := ""

// 	if num == 0 {
// 		return "0"
// 	}

// 	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
// 	for ; num > 0; num /= 2 {
// 		bin := num % 2
// 		// 注意这里不能写成 s += strconv.Itoa(bin)，否则二进制顺序颠倒
// 		s = strconv.Itoa(bin) + s
// 	}

// 	return s
// }

// // countOne 计算1的个数
// func countOne(s string) int {

// 	count := 0
// 	for _, v := range s {
// 		// 字符串遍历出是int32
// 		// rune 等同于int32,常用来处理unicode或utf-8字符
// 		// byte 等同于int8，常用来处理ascii字符
// 		if string(v) == "1" {
// 			count += 1
// 		}
// 	}

// 	return count
// }

// @lc code=end

