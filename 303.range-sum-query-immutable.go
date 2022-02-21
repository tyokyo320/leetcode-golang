/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 */

// @lc code=start
/*
	type NumArray struct {
		nums []int
	}

	func Constructor(nums []int) NumArray {
		return NumArray{nums}
	}

	func (this *NumArray) SumRange(left int, right int) int {
		var sumRange int = 0

		for i := left; i <= right; i++ {
			sumRange += this.nums[i]
		}

		return sumRange
	}
*/

type NumArray struct {
	sums []int
}

// Constructor 在初始化的时候计算出数组 nums 在每个下标处的前缀和
func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, v := range nums {
		sums[i+1] = sums[i] + v
	}
	return NumArray{sums}
}

func (na *NumArray) SumRange(i, j int) int {
	return na.sums[j+1] - na.sums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

