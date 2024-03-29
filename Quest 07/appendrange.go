package piscine

func AppendRange(min, max int) []int {
	ns := []int{}
	if min >= max {
		return nil
	}
	for i := min; i < max; i++ {
		ns = append(ns, i)
	}
	return ns
}
