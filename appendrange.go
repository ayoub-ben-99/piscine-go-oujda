package piscine

func AppendRange(min, max int) []int {
	rangeInt := []int{}
	if min >= max {
		return nil
	}
	for i := min; i < max; i++ {
		rangeInt = append(rangeInt, i)
	}
	return rangeInt
}
