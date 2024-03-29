package piscine

import "github.com/01-edu/z01"

func PrintR(r rune) {
	z01.PrintRune(r)
}

func DescendComb() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for k := '9'; k >= '0'; k-- {
				for l := '9'; l >= '0'; l-- {
					if (i-48)*10+j > (k-48)*10+l {
						PrintR(i)
						PrintR(j)
						PrintR(' ')
						PrintR(k)
						PrintR(l)
						if i == '0' && j == '1' && k == '0' && l == '0' {
							break
						}
						PrintR(',')
						PrintR(' ')
					}
				}
			}
		}
	}
}
