package piscine

func BasicAtoi(s string) int {
	n := 0
	a := 0
	for i := range s {
		if s[i] == '0' {
			a++
			continue
		}
		break
	}
	for i := a; i < len(s); i++ {
		n = ((n * 10) + int(s[i]) - 48)
	}
	return n
}
