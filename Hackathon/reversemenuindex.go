package piscine

func ReverseMenuIndex(menu []string) []string {
	ns := make([]string, len(menu))
	c := 0
	for i := len(menu) - 1; i >= 0; i-- {
		ns[c] = menu[i]
		c++
	}
	return ns
}
