package piscine

func f(a, b int) int {
	if a == b {
		return 0
	}
	return a - b
}

func IsSorted(f func(a, b int) int, a []int) bool {
	ac := true
	dc := true
	for i := 0; i < len(a)-2; i++ {
		if f(a[i], a[i+1]) > 0 {
			ac = false
		}
	}
	for i := 0; i < len(a)-2; i++ {
		if f(a[i], a[i+1]) < 0 {
			dc = false
		}
	}
	if ac {
		return ac
	}
	return dc
}
