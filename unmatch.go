package piscine

func Unmatch(a []int) int {
	for i := range a {
		count := 0
		for j := range a {
			if a[i] == a[j] {
				count++
			}
		}
		if count%2 != 0 {
			return a[i]
		}
	}
	return -1
}
