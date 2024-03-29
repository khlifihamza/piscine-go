package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		for _, v := range args[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
