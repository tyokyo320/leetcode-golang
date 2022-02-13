/*
 * @lc app=leetcode id=976 lang=golang
 *
 * [976] Largest Perimeter Triangle
 */

// @lc code=start
func largestPerimeter(nums []int) int {

	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	length := len(nums)

	// 假设a>b>c的话，我们实际上只需要判断b+c>a即可，也就是判断最小的两个值的和是不是最大值即可
	for i := 0; i < length-2; i++ {
		if nums[i+2]+nums[i+1] > nums[i] {
			return nums[i+2] + nums[i+1] + nums[i]
		}
	}

	return 0
}

// @lc code=end
