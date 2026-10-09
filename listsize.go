package piscine

func ListSize(l *List) int {
	count := 0
	fst := l.Head
	for fst != nil {
		count++
		fst = fst.Next
	}
	return count
}
