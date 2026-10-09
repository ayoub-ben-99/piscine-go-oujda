package piscine

func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n"
	}
	result := ""
	for i := 2; i < len(str); i = i + 3 {
		result += string(rune(str[i]))
	}
	return result + "\n"
}
