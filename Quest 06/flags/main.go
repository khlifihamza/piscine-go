package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func Contains(s string, toFind string) bool {
	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("\t This flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
	} else {
		if args[0] == "--help" || args[0] == "-h" {
			fmt.Println("--insert")
			fmt.Println("  -i")
			fmt.Println("\t This flag inserts the string into the string passed as argument.")
			fmt.Println("--order")
			fmt.Println("  -o")
			fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
		} else {
			ns := []string{}
			s := ""
			o := false
			rev := false
			for _, v := range args {
				s = ""
				if Contains(v, "--insert") {
					for j := len("--insert"); j < len(v)-1; j++ {
						s += string(v[j+1])
					}
					ns = append(ns, s)
					continue
				} else if Contains(v, "--order") {
					for j := len("--order"); j < len(v)-1; j++ {
						s += string(v[j+1])
					}
					ns = append(ns, s)
					o = true
					continue
				} else if Contains(v, "-i") {
					for j := len("-i"); j < len(v)-1; j++ {
						s += string(v[j+1])
					}
					ns = append(ns, s)
					continue
				} else if Contains(v, "-o") {
					for j := len("-o"); j < len(v)-1; j++ {
						s += string(v[j+1])
					}
					ns = append(ns, s)
					o = true
					continue
				} else {
					rev = true
					ns = append(ns, v)
					continue
				}
			}
			if o {
				rs := []rune{}
				var t rune
				for _, v := range ns {
					for _, x := range v {
						rs = append(rs, x)
					}
				}
				for i := 0; i < len(rs)-1; i++ {
					for j := i + 1; j < len(rs); j++ {
						if rs[i] > rs[j] {
							t = rs[i]
							rs[i] = rs[j]
							rs[j] = t
						}
					}
				}
				for _, v := range rs {
					z01.PrintRune(v)
				}
				z01.PrintRune('\n')
			} else if rev {
				for i := len(ns) - 1; i >= 0; i-- {
					for _, x := range ns[i] {
						z01.PrintRune(x)
					}
				}
				z01.PrintRune('\n')
			} else {
				for _, v := range ns {
					for _, x := range v {
						z01.PrintRune(x)
					}
				}
				z01.PrintRune('\n')
			}
		}
	}
}
