package piscine

func Index(s string, toFind string) int {
	if len(s) == 0 || len(toFind) > len(s) {
		return -1
	}
	c := 0

	index := 0
	if len(toFind) == 1 {
		for i := 0; i < len(s); i++ {
			if s[i] == toFind[c] {
				index = i
				break
			}
		}
	}
	if len(toFind) > 1 {
		for i := 0; i < len(s); i++ {
			if s[i] == toFind[c] {
				index = i
				break
			}
		}
		c = index
		for j := 0; j < len(toFind); j++ {
			if s[c] == toFind[j] {
				c++
			} else {
				index = -1
			}
		}
	}
	return index
}
