package piscine

func IsAlpha(s string) bool {
	alphanum := false
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9') || (s[i] >= 'a' && s[i] <= 'z') {
			alphanum = true
		} else {
			return false
		}
	}
	return alphanum
}
