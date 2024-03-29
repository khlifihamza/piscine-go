package piscine

func Compare(a, b string) int {
	var n int
	r := 0
	if len(a) < len(b) {
		n = len(a)
	} else {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		if a[i] == b[i] {
			r = 0
		}
		if rune(a[i]) < rune(b[i]) || len(a) > len(b) {
			r = 1
		}
		if rune(a[i]) > rune(b[i]) || len(a) < len(b) {
			r = -1
		}
	}
	return r
}
