package piscine

func Split(s, sep string) []string {
	count := 1

	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			count++
		}
	}

	result := make([]string, count)

	start := 0
	index := 0

	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			result[index] = s[start:i]
			index++
			start = i + len(sep)
		}
	}

	result[index] = s[start:]

	return result
}
