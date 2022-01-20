/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })
	sort.Slice(nums2, func(i, j int) bool { return nums2[i] < nums2[j] })
	length1 := len(nums1)
	length2 := len(nums2)
	result := []int{}
	i := 0
	j := 0

	for i < length1 && j < length2 {
		if nums1[i] < nums2[j] {
			i += 1
		} else if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i += 1
			j += 1
		} else {
			j += 1
		}
	}
	return result
}

// @lc code=end

