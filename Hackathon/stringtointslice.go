package piscine

func StringToIntSlice(str string) []int {
	ns := []int{}
	if str == "" {
		return nil
	}
	for _, v := range str {
		ns = append(ns, int(v))
	}
	return ns
}
