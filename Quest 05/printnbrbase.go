package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	nbase := []rune(base)
	nv := false
	ns := []int{}
	for i := 0; i < len(base)-1; i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				nv = true
				break
			}
		}
	}
	for _, v := range nbase {
		if v == '+' || v == '-' {
			nv = true
			break
		}
	}
	if nv || len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		nbr = 223372036854775808
	}
	if nv == false {
		basenum := len(base)
		mod := 0
		if nbr < 0 {
			nbr *= -1
			z01.PrintRune('-')
		}
		for nbr > 0 {
			mod = nbr % basenum
			ns = append(ns, mod)
			nbr /= basenum
		}
		for i := len(ns) - 1; i >= 0; i-- {
			z01.PrintRune(rune(base[ns[i]]))
		}
	}
}
