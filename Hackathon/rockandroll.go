package piscine

func RockAndRoll(n int) string {
	s := ""
	if n < 0 {
		s = "error: number is negative\n"
	} else if n%2 == 0 && n%3 == 0 {
		s = "rock and roll\n"
	} else if n%2 == 0 {
		s = "rock\n"
	} else if n%3 == 0 {
		s = "roll\n"
	} else {
		s = "error: non divisible\n"
	}
	return s
}
