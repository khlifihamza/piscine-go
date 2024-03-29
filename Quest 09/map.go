package piscine

func Map(f func(int) bool, a []int) []bool {
	bs := []bool{}
	for _, v := range a {
		if f(v) {
			bs = append(bs, true)
		} else {
			bs = append(bs, false)
		}
	}
	return bs
}
