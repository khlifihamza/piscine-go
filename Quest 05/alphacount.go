package piscine

func AlphaCount(s string) int {
	c := 0
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			c++
		}
	}
	return c
}
