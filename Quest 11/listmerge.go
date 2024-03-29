package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
		return
	} else if l2.Head == nil {
		l2.Head = l1.Head
		return
	} else {
		l1.Tail.Next = l2.Head
	}
}
