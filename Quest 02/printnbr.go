package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var min int = -9223372036854775808
	if n == min {
		var strmin string = "-9223372036854775808"
		for i := 0; i < len(strmin); i++ {
			z01.PrintRune(rune(strmin[i]))
		}
	} else if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 0 && n < 10 {
		z01.PrintRune(rune(n + '0'))
	}
	if n >= 10 {
		PrintNbr(n / 10)
		PrintNbr(n % 10)
	}
}
