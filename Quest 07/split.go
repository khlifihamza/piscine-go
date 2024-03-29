package piscine

func Split(s, sep string) []string {
	ns := []string{}
	ap := ""
	in := 0
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			for j := in; j < i; j++ {
				ap += string(s[j])
			}
			ns = append(ns, ap)
			in = i + len(sep)
			if i < len(s) {
				ap = ""
			}
		}
	}
	for i := in; i < len(s); i++ {
		ap += string(s[i])
	}
	ns = append(ns, ap)
	return ns
}
