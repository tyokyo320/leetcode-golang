/*
 * @lc app=leetcode id=1281 lang=golang
 *
 * [1281] Subtract the Product and Sum of Digits of an Integer
 */

// @lc code=start
func subtractProductAndSum(n int) int {

	/*
		使用strconv包

		// strconv包提供了字符串与简单数据类型之间的类型转换功能
		// 字符串转int：Atoi()
		// int转字符串: Itoa()
		length := len(strconv.Itoa(n))

		productOfDigits := 1
		sumOfDigits := 0
		num := 0

		// eg.234: 2 -> 3 -> 4
		for i := length - 1; i >= 0; i-- {
			num = n / int(math.Pow10(i))
			productOfDigits *= num
			sumOfDigits += num
			n = n - num*int(math.Pow10(i))
		}

		return productOfDigits - sumOfDigits
	*/

	// eg.234: 4 -> 3 -> 2
	sum, product := 0, 1
	for n > 0 {
		mod := n % 10
		sum += mod
		product *= mod

		n /= 10

	}

	return product - sum
}

// @lc code=end

