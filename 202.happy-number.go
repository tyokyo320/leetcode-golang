/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {

	/*
		有2种情况：第一种经过变换最终为1；第二种，经过变换会出现相同的数字，导致循环
		和“循环”有关的问题可以考虑一下“快慢指针”
	*/
	slow, quick := n, n

	for {
		quick = transform(transform(quick))
		slow = transform(slow)
		if quick == 1 {
			return true
		}
		if quick == slow {
			return false
		}
	}
}

// transform 求平方和
func transform(n int) int {
	var sum int = 0

	for n > 0 {
		sum += int(math.Pow(float64(n%10), float64(2)))
		n /= 10
	}

	return sum
}

// @lc code=end

