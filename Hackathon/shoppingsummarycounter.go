package piscine

func SplitS(s, sep string) []string {
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

func ShoppingSummaryCounter(str string) map[string]int {
	summap := map[string]int{}
	ns := SplitS(str, " ")
	for i := 0; i < len(ns); i++ {
		c := 0
		for j := 0; j < len(ns); j++ {
			if ns[i] == ns[j] {
				c++
			}
		}
		summap[ns[i]] = c
	}
	return summap
}
