package piscine

func IsNumeric(s string) bool {
	num := false
	for _, v := range s {
		if v >= '0' && v <= '9' {
			num = true
		} else {
			return false
		}
	}
	return num
}
