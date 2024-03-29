package piscine

func IterativePower(nb int, power int) int {
	if power >= 1 {
		n := nb
		for i := 0; i < power-1; i++ {
			nb = nb * n
		}
	} else if power == 0 {
		return 1
	} else {
		return 0
	}
	return nb
}
