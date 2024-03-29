package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	ns := []int{}
	for n > 0 {
		ns = append(ns, n%10)
		n /= 10
	}
	t := 0
	for i := 0; i < len(ns)-1; i++ {
		for j := 0; j < len(ns)-1-i; j++ {
			if ns[j] > ns[j+1] {
				t = ns[j+1]
				ns[j+1] = ns[j]
				ns[j] = t
			}
		}
	}
	for _, v := range ns {
		z01.PrintRune(rune(v) + '0')
	}
}
