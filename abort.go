package piscine

func Abort(a, b, c, d, e int) int {
	numbers := []int{a, b, c, d, e}
	for i := 1; i < len(numbers); i++ {
		current := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > current {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = current
	}
	return numbers[2]
}
