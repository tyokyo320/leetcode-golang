/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	var stack []byte

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) == 0 {
			return false
		} else if stack[len(stack)-1] == '(' && s[i] == ')' {
			stack = stack[:len(stack)-1]
		} else if stack[len(stack)-1] == '{' && s[i] == '}' {
			stack = stack[:len(stack)-1]
		} else if stack[len(stack)-1] == '[' && s[i] == ']' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

// @lc code=end

