package piscine

func isprime(nb int) bool {
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	for i := 5; i*i <= nb; i = i + 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	if nb == 3 {
		return 3
	}
	for i := 2; i <= nb/2; i++ {
		if isprime(nb) {
			return nb
		} else {
			nb++
		}
	}
	return nb
}
