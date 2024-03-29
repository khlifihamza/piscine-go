package piscine

func AtoiBase(s string, base string) int {
	for i := 0; i < len(base)-1; i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}
	for _, v := range base {
		if v == '+' || v == '-' {
			return 0
		}
	}
	if len(base) < 2 {
		return 0
	}
	n := 0
	for _, v := range s {
		nin := true
		in := 0
		for i, x := range base {
			if v == x {
				in = i
				nin = false
			}
		}
		if nin {
			return 0
		}
		n = n*len(base) + in
	}
	return n
}
