package piscine

func ToLower(s string) string {
	ns := []rune(s)
	for i, v := range ns {
		if v >= 'A' && v <= 'Z' {
			ns[i] += 32
		}
	}
	return string(ns)
}
