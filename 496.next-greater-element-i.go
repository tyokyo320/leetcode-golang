/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	/*
		length1 := len(nums1)
		length2 := len(nums2)
		result := []int{}

		i, pointer, counter, temp := 0, 0, 0, 0
		for i < length1 {
			for j, _ := range nums2 {
				if nums1[i] == nums2[j] {
					pointer = j + 1
					temp = j
					if j == length2-1 {
						// fmt.Printf("i = %v, temp = %v\n", i, temp)
						result = append(result, -1)
					}
				}
			}

			counter = 0
			for pointer < length2 {
				if nums2[pointer] > nums1[i] {
					result = append(result, nums2[pointer])
					break
				} else {
					counter += 1
					// fmt.Printf("i = %v, temp = %v, pointer = %v, counter = %v\n", i, temp, pointer, counter)
					if pointer == length2-1 && counter == length2-1-temp {
						// fmt.Printf("i = %v, temp = %v\n", i, temp)
						result = append(result, -1)
					}
					pointer += 1
				}
			}
			i += 1
		}

		return result
	*/

	/*
		单调栈 + 哈希表
		1. 由于数组中不存在重复元素，因此我们使用字典（哈希表）来保存每个数对应的下一个更大元素。
		首先遍历nums2，完成哈希表的建立，然后对nums1的每个数，查表即可
		2. 完成哈希表的过程，我们使用单调栈，
		以示例1来说，nums2 = [1,3,4,2]，我们首先将1入栈，然后是3，3与栈最后一个元素对比，大于它，那么1出栈，1对应的答案即为3，
		将3入栈，下一个是4，4大于目前栈的最后一个元素3，所以3出栈，3的答案就是4，
		将4入栈，下一个是2，2小于目前栈的最后一个元素4，所以直接将2入栈。
		目前栈内剩余的[4,2]，两个元素右边不存在比自己大的元素，所以它们在哈希表中对应的答案为-1。

		-> 根据上面的过程，我们的栈维护的是一个递减的数列，对于新入栈的元素，把小于这个元素的栈内元素pop，
		直到结束后，这个栈内的所有元素是递减的，也就肯定不存在右边第一个更大元素了，它们哈希表中对应的为-1。
	*/

	// 单调栈
	stack := []int{}
	dict := make(map[int]int)
	length2 := len(nums2)
	remove := 0

	for i := 0; i < length2; i++ {
		// 遍历nums2，入栈，将栈内小于它的元素出栈，填写dict: 保存每个数对应的下一个更大元素
		for len(stack) > 0 && (nums2[i] > stack[len(stack)-1]) {
			// 获取列表最后一个元素
			remove = stack[len(stack)-1]
			// 删除列表中最后一个元素（stack.pop()）
			stack = stack[:len(stack)-1]
			dict[remove] = nums2[i]
		}
		stack = append(stack, nums2[i])
	}

	// 最后stack里面剩下的元素，都是右边没有比它大的元素，所以它们对应-1
	for len(stack) != 0 {
		remove = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		dict[remove] = -1
	}

	result := []int{}
	// 对于num1中直接查表即可
	for _, num := range nums1 {
		result = append(result, dict[num])
	}

	return result
}

// @lc code=end

