package main

import (
	"fmt"
)

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}

	fmt.Print(nil, "\n")
}

func main() {
	link := &List{}
	link2 := &List{}

	fmt.Println("----normal state----")
	ListPushBack(link2, 1)
	PrintList(link2)
	ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	ListPushBack(link, 1)
	ListPushBack(link, "Hello")
	ListPushBack(link, 1)
	ListPushBack(link, "There")
	ListPushBack(link, 1)
	ListPushBack(link, 1)
	ListPushBack(link, "How")
	ListPushBack(link, 1)
	ListPushBack(link, "are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)
	PrintList(link)

	ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

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

func ListRemoveIf(l *List, data_ref interface{}) {
	h := l.Head
	l.Head = nil
	l.Tail = nil
	if h.Data != data_ref {
		ListPushBack(l, h.Data)
	}
	for h.Next != nil {
		if h.Next.Data == data_ref {
			if h.Next.Data == data_ref {
				h = h.Next
			}
		} else {
			ListPushBack(l, h.Next.Data)
			h = h.Next
		}
	}
}
