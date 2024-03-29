package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadD() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	x, _ := strconv.Atoi(args[0])
	y, _ := strconv.Atoi(args[1])
	if x > 0 && y > 0 {
		for j := 1; j <= y; j++ {
			for i := 1; i <= x; i++ {
				if j == 1 || j == y {
					if i == 1 {
						z01.PrintRune('A')
					} else if i == x {
						z01.PrintRune('C')
					} else {
						z01.PrintRune('B')
					}
				} else {
					if i == 1 || i == x {
						z01.PrintRune('B')
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune('\n')
		}
	}
}

func main() {
	QuadD()
}
