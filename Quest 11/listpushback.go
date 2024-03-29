package piscine

func ListPushBack(l *List, data interface{}) {
	h := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = h
		l.Tail = h
	} else {
		l.Tail.Next = h
		l.Tail = h
	}
}
