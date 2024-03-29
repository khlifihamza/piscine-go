package piscine

func ListPushFront(l *List, data interface{}) {
	h := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = h
		l.Tail = h
	} else {
		h.Next = l.Head
		l.Head = h
	}
}
