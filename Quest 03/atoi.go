package piscine

func Atoi(s string) int {
	n := 0
	a := 0
	neg := 0
	pos := 0
	for i := range s {
		if s[i] == '-' {
			neg++
			a++
			continue
		}
		break
	}
	for i := range s {
		if s[i] == '+' {
			pos++
			a++
			continue
		}
		break
	}

	for i := range s {
		if s[i] == '0' {
			a++
			continue
		}
		break
	}

	for i := a; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			n = ((n * 10) + int(s[i]) - 48)
		} else {
			n = 0
			break
		}
	}
	if neg == 1 {
		n = n * -1
	}
	if neg > 1 {
		n = 0
	}
	if pos == 1 {
		n = n
	}
	if pos > 1 {
		n = 0
	}
	return n
}
