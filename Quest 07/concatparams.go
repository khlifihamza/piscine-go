package piscine

func ConcatParams(args []string) string {
	s := ""
	for i, v := range args {
		s += v
		if i < len(args)-1 {
			s += "\n"
		}
	}
	return s
}
