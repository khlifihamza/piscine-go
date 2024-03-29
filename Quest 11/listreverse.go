package piscine

func ListReverse(l *List) {
	reserve := l.Head
	cur := l.Head
	var prev *NodeL
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}
	l.Head = prev
	l.Tail = reserve
}
