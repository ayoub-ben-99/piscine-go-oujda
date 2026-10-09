package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}
	p := 0
	for l != nil {
		if p == pos {
			return l
		}
		l = l.Next
		p++
	}
	return nil
}
