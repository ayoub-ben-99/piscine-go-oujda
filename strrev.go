package piscine

func StrRev(s string) string {
	newStr := ""
	for i := 0; i < len(s); i++ {
		newStr += string(s[len(s)-1-i])
	}
	return newStr
}
