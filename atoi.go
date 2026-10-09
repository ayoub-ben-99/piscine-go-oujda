package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0
	switch s[0] {
	case '-':
		sign = -1
		start = 1
	case '+':
		start = 1
	}

	n := 0
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		n = n*10 + int(s[i]-'0')
	}

	return n * sign
}
