package piscine

func IsPrintable(s string) bool {
	print := false
	for _, v := range s {
		if v > 31 {
			print = true
		} else {
			return false
		}
	}
	return print
}
