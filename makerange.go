package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	rangeInt := make([]int, max-min)
	for i := min; i < max; i++ {
		rangeInt[i-min] = i
	}
	return rangeInt
}
