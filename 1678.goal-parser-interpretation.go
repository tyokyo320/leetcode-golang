/*
 * @lc app=leetcode id=1678 lang=golang
 *
 * [1678] Goal Parser Interpretation
 */

// @lc code=start
func interpret(command string) string {
	length := len(command)
	result := ""
	i := 0

	for i < length {
		if string(command[i]) == "G" {
			result += "G"
			i += 1
		} else if (string(command[i]) == "(") && (string(command[i+1]) == ")") {
			result += "o"
			i += 2
		} else if (string(command[i]) == "(") && (string(command[i+1]) == "a") {
			result += "al"
			i += 4
		}
	}

	return result
}

// @lc code=end

