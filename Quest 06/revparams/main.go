package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := len(args) - 1; i >= 1; i-- {
		for _, v := range args[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
