package piscine

func Max(a []int) int {
	n := 0
	if len(a) == 0 {
		return 0
	}
	for i := 0; i < len(a); i++ {
		if a[i] > n {
			n = a[i]
		}
	}
	return n
}
