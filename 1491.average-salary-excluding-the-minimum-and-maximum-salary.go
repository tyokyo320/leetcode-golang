/*
 * @lc app=leetcode id=1491 lang=golang
 *
 * [1491] Average Salary Excluding the Minimum and Maximum Salary
 */

// @lc code=start
func average(salary []int) float64 {
	sum := 0
	length := len(salary)

	// 排序方法
	// sort.Ints(salary)

	// 逆序排序
	// ints := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	// sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	// fmt.Println(ints) // [9 8 7 6 5 4 3 2 1 0]
	sort.Slice(salary, func(i, j int) bool { return salary[i] < salary[j] })

	for i := 1; i < length-1; i++ {
		sum += salary[i]
	}

	return float64(sum) / float64(length-2)
}

// @lc code=end

