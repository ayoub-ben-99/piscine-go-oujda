package piscine

func Compact(ptr *[]string) int {
	count := 0
	newArr := *ptr
	*ptr = []string{}
	for _, v := range newArr {
		if v != "" {
			*ptr = append(*ptr, v)
			count++
		}
	}
	return count
}
