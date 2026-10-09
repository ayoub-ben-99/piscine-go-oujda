package piscine

func IterativeFactorial(nb int) int {
	factorial := 1
	if nb == 0 || nb == 1 {
		return 1
	}
	if nb < 0 {
		return 0
	}
	for i := 2; i <= nb; i++ {
		next := factorial * i
		if next/i != factorial {
			return 0
		}
		factorial = next
	}
	return factorial
}
