package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	n := 0
	ns := []rune{}
	if len(args) == 1 {
		return
	}
	if args[1] == "--upper" {
		for i := 2; i < len(args); i++ {
			if i >= 1 {
				if len(args[i]) == 1 && args[i] >= "0" && args[i] <= "9" {
					ns = append(ns, rune(args[i][0])+16)
				} else if len(args[i]) >= 2 {
					if args[i] >= "10" && args[i] <= "26" {
						n = 0
						for j := 0; j < len(args[i]); j++ {
							n = ((n * 10) + int(args[i][j]) - 48)
						}
						ns = append(ns, rune(n+64))
					} else {
						ns = append(ns, rune(32))
					}
				} else {
					ns = append(ns, rune(32))
				}
			}
		}
		for _, v := range ns {
			z01.PrintRune(v)
		}
	} else {
		for i, v := range args {
			if i >= 1 {
				if len(v) == 1 && v >= "0" && v <= "9" {
					ns = append(ns, rune(v[0])+48)
				} else if len(v) >= 2 {
					if v >= "10" && v <= "26" {
						n = 0
						for i := 0; i < len(v); i++ {
							n = ((n * 10) + int(v[i]) - 48)
						}
						ns = append(ns, rune(n+96))
					} else {
						ns = append(ns, rune(32))
					}
				} else {
					ns = append(ns, rune(32))
				}
			}
		}
		for _, v := range ns {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
