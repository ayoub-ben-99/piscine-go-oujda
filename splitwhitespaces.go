package piscine

func SplitWhiteSpaces(str string) []string {
	result := []string{}
	word := ""

	for _, r := range str {
		if r == ' ' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}

	if word != "" {
		result = append(result, word)
	}

	return result
}
