package piscine

func Rot14(s string) string {
	ns := ""
	n := '0'
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			if v+14 <= 'z' {
				ns += string(v + 14)
			} else {
				n = 'z' - v
				ns += string(13 - n + 'a')
			}
		}
		if v >= 'A' && v <= 'Z' {
			if v+14 <= 'Z' {
				ns += string(v + 14)
			} else {
				n = 'Z' - v
				ns += string(13 - n + 'A')
			}
		}
		if (v < 'a' || v > 'z') && (v < 'A' || v > 'Z') {
			ns += string(v)
		}
	}
	return ns
}
