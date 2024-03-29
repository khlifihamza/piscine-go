package piscine

func TrimAtoi(s string) int {
	ns := []int{}
	neg := 1
	a := 0
	fn := false
	for i := range s {
		if s[i] == '0' {
			a++
			continue
		}
		break
	}
	for i := a; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			ns = append(ns, int(s[i]))
			fn = true
		}
		if s[i] == '-' && fn == false {
			neg = -1
		}
	}
	res := 0
	for i := 0; i < len(ns); i++ {
		res = ((res * 10) + int(ns[i]) - 48)
	}
	return neg * res
}
