package piscine

func SplitWhiteSpaces(s string) []string {
	ns := []string{}
	ap := ""
	for i, v := range s {
		if v == '\t' || v == '\n' || v == ' ' {
			if string(s[i-1]) == " " {
				continue
			}
			ns = append(ns, ap)
			if i < len(s)-1 {
				ap = ""
				continue
			}
		}
		ap += string(v)
	}
	ns = append(ns, ap)
	return ns
}
