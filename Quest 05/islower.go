package piscine

func IsLower(s string) bool {
	low := false
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			low = true
		} else {
			return false
		}
	}
	return low
}
