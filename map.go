package piscine

func Map(f func(int) bool, a []int) []bool {
	arr := []bool{}
	for _, v := range a {
		if !f(v) {
			arr = append(arr, false)
		} else {
			arr = append(arr, true)
		}
	}
	return arr
}
