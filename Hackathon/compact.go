package piscine

func Compact(ptr *[]string) int {
	c := 0
	ns := []string{}
	for _, v := range *ptr {
		if v != "" {
			c++
			ns = append(ns, v)
		}
	}
	*ptr = ns
	return c
}
