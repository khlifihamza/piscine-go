package main

import (
	"github.com/01-edu/z01"
)

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	z01.PrintRune('\n')
	ptrDoor.state = "CLOSE"
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	z01.PrintRune('\n')
	ptrDoor.state = "OPEN"
}

func IsDoorOpen(Door Door) bool {
	PrintStr("is the Door opened ?")
	z01.PrintRune('\n')
	if Door.state == "OPEN" {
		return true
	}
	return false
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	z01.PrintRune('\n')
	if ptrDoor.state == "CLOSE" {
		return true
	}
	return false
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(*door) {
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}
}
