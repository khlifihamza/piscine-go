package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintR(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

const s string = "x = 42, y = 21\n"

func main() {
	points := &point{}

	setPoint(points)

	PrintR(s)
}
