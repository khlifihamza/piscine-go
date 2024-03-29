package piscine

func JumpOver(str string) string {
	s := ""
	c := 2
	if len(str) < 3 {
		return "\n"
	}
	for i := 1; i < len(str); i++ {
		if i == c {
			s += string(str[i])
			c += 3
		}
	}
	s += "\n"
	return s
}
