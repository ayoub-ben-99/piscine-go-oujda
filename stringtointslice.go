package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}
	newSlice := []int{}
	for _, char := range str {
		newSlice = append(newSlice, int(char))
	}
	return newSlice
}
