package piscine

func Unmatch(a []int) int {
	n := -1
	c := 0
	for _, v := range a {
		for _, x := range a {
			if v == x {
				c++
			}
		}
		if c%2 != 0 {
			n = v
			break
		}
	}
	return n
}
