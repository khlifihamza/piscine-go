package piscine

func BasicAtoi2(s string) int {
	a := 0
	n := 0
	for i := range s {
		if s[i] == '0' {
			a++
			continue
		}
		break
	}
	for i := a; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			n = ((int(s[i]) + n*10) - 48)
		} else {
			n = 0
			break
		}
	}
	return n
}
