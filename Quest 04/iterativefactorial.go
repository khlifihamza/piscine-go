package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb >= 1 && nb <= 20 {
		fac := 1
		for i := 1; i <= nb; i++ {
			fac = fac * i
		}
		return fac
	} else {
		return 0
	}
}
