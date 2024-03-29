package piscine

func IsUpper(s string) bool {
	up := false
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			up = true
		} else {
			return false
		}
	}
	return up
}
