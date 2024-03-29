package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for _, v := range args[0] {
		if v == '.' || v == '/' {
			continue
		} else {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
