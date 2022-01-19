/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {

	// 双指针 + 逆序赋值
	i := m - 1
	j := n - 1
	index := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[index] = nums1[i]
			i -= 1
		} else {
			nums1[index] = nums2[j]
			j -= 1
		}
		index -= 1
	}

	// 当nums1或nums2遍历结束，另一数组有剩余时，将剩余数组赋予剩余位置
	for i >= 0 {
		nums1[index] = nums1[i]
		index -= 1
		i -= 1
	}

	for j >= 0 {
		nums1[index] = nums2[j]
		index -= 1
		j -= 1
	}

}

// @lc code=end

