package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	t := ""
	for i := 1; i < len(args); i++ {
		for j := 1; j < len(args)-i; j++ {
			if args[j] > args[j+1] {
				t = args[j+1]
				args[j+1] = args[j]
				args[j] = t
			}
		}
	}
	for i := range args {
		for _, v := range args[i] {
			z01.PrintRune(rune(v))
		}
		z01.PrintRune('\n')
	}
}
