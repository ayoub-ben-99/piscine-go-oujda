package piscine

func ReverseMenuIndex(menu []string) []string {
	newSlice := make([]string, len(menu))

	for i := len(menu) - 1; i >= 0; i-- {
		newSlice[len(menu)-1-i] = menu[i]
	}
	return newSlice
}
