package piscine

func Capitalize(s string) string {
	fl := false
	ns := []rune(s)
	for i := 0; i < len(s); i++ {
		if (ns[i] >= 'a' && ns[i] <= 'z') || (ns[i] >= 'A' && ns[i] <= 'Z') {
			if ns[i] >= 'a' && ns[i] <= 'z' && fl == false {
				ns[i] -= 32
				fl = true
				continue
			} else if ns[i] >= 'A' && ns[i] <= 'Z' && fl == false {
				fl = true
				continue
			} else if ns[i] >= 'A' && ns[i] <= 'Z' && fl == true {
				ns[i] += 32
			}
		}
		if (ns[i] >= 'A' && ns[i] <= 'Z') || (ns[i] >= '0' && ns[i] <= '9') || (ns[i] >= 'a' && ns[i] <= 'z') {
			fl = true
		} else {
			fl = false
		}
	}
	return string(ns)
}
