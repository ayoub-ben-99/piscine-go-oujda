package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)

	newStr := ""

	for _, char := range str {
		if char == ' ' {
			result[newStr]++
			newStr = ""
		} else {
			newStr += string(char)
		}
	}

	result[newStr]++

	return result
}
