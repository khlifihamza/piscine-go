package piscine

func BasicJoin(elems []string) string {
	ns := []rune{}
	for i := 0; i < len(elems); i++ {
		for _, v := range elems[i] {
			ns = append(ns, v)
		}
	}
	return string(ns)
}
