/*
 * @lc app=leetcode id=1309 lang=golang
 *
 * [1309] Decrypt String from Alphabet to Integer Mapping
 */

// @lc code=start
func freqAlphabets(s string) string {
	/*
		补充2个空格以应对越界问题

		length := len(s)
		result := ""
		s += "  "
		i := 0

		for i < length {
			if string(s[i+2]) == "#" {
				result += mapping(s[i : i+3])
				i += 3
			} else {
				result += mapping(s[i : i+1])
				i += 1
			}
		}

		return result
	*/

	length := len(s)
	result := ""
	i := 0

	for i < length {
		if (i+2 < length) && (string(s[i+2]) == "#") {
			result += mapping(s[i : i+3])
			i += 3
		} else {
			result += mapping(s[i : i+1])
			i += 1
		}
	}

	return result
}

func mapping(key string) string {
	var dict = map[string]string{
		"1":   "a",
		"2":   "b",
		"3":   "c",
		"4":   "d",
		"5":   "e",
		"6":   "f",
		"7":   "g",
		"8":   "h",
		"9":   "i",
		"10#": "j",
		"11#": "k",
		"12#": "l",
		"13#": "m",
		"14#": "n",
		"15#": "o",
		"16#": "p",
		"17#": "q",
		"18#": "r",
		"19#": "s",
		"20#": "t",
		"21#": "u",
		"22#": "v",
		"23#": "w",
		"24#": "x",
		"25#": "y",
		"26#": "z",
	}

	return dict[key]
}

// @lc code=end

