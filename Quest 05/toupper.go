package piscine

func ToUpper(s string) string {
	ns := []rune(s)
	for i, v := range ns {
		if v >= 'a' && v <= 'z' {
			ns[i] -= 32
		}
	}
	return string(ns)
}
