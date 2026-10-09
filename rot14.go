package piscine

func Rot14(s string) string {
	newStr := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			newStr += string('a' + (c-'a'+14)%26)
		} else if c >= 'A' && c <= 'Z' {
			newStr += string('A' + (c-'A'+14)%26)
		} else {
			newStr += string(c)
		}
	}
	return newStr
}
