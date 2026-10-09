package piscine

func DescendAppendRange(max, min int) []int {
	newArr := []int{}
	for i := max; i > min; i-- {
		newArr = append(newArr, i)
	}
	return newArr
}
