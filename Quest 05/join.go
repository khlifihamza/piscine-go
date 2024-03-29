package piscine

func Join(strs []string, sep string) string {
	ns := []rune{}
	nsep := []rune(sep)
	for i := 0; i < len(strs); i++ {
		for _, v := range strs[i] {
			ns = append(ns, v)
		}
		if i < len(strs)-1 {
			ns = append(ns, nsep...)
		}
	}
	return string(ns)
}
