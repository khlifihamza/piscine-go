package piscine

func CompareS(a, b string) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if CompareS(a[j], a[j+1]) > 0 {
				t := a[j+1]
				a[j+1] = a[j]
				a[j] = t
			}
		}
	}
}
