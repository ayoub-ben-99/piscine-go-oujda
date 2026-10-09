package piscine

func FindNextPrime(nb int) int {
	count := nb
	for !IsPrime(count) {
		count++
	}
	return count
}
