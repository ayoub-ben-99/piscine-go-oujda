package piscine

func ListPushBack(l *List, data interface{}) {
	b := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = b
		l.Tail = b
	} else {
		l.Tail.Next = b
		l.Tail = b
	}
}
