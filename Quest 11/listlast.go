package piscine

func ListLast(l *List) interface{} {
	var d interface{}
	for l.Head != nil {
		if l.Head.Next == nil {
			d = l.Head.Data
		}
		l.Head = l.Head.Next
	}
	return d
}
