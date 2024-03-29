package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	var p *NodeL
	c := 0
	for l != nil {
		if c == pos {
			p = l
		}
		c++
		l = l.Next
	}
	return p
}
