package quick_func

import "strconv"

func SubString(s string, start, length int) string {
	if start < 0 {
		start = 0
	}
	sLen := len(s)

	if start > sLen {
		return ""
	}

	t := len(s[start:sLen])
	if t < length {
		length = t
	}

	if length < 0 {
		length = t + length
	}

	return s[start : start+length]
}

func StringToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
