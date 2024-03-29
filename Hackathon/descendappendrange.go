package piscine

func DescendAppendRange(max, min int) []int {
	ns := []int{}
	if max <= min {
		return ns
	}
	for i := max; i > min; i-- {
		ns = append(ns, i)
	}
	return ns
}
