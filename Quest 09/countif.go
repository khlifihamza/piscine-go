package piscine

func CountIf(f func(string) bool, tab []string) int {
	c := 0
	for _, v := range tab {
		if f(v) {
			c++
		}
	}
	return c
}
