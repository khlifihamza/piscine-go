package piscine

func NRune(s string, n int) rune {
	if n < 1 || len(s) < n {
		return 0
	}
	ns := []rune(s)
	return ns[n-1]
}
