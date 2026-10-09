package piscine

func ListReverse(l *List) {
	var pr, nxt *NodeL
	cr := l.Head
	for cr != nil {
		nxt = cr.Next
		cr.Next = pr
		pr = cr
		cr = nxt
	}
	l.Head = pr
}
