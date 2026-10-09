package piscine

func ListLast(l *List) interface{} {
	if l == nil || l.Tail == nil {
		return nil
	} else {
		return l.Tail.Data
	}
}
