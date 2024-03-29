package piscine

func SortIntegerTable(table []int) {
	t := 0
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1-i; j++ {
			if table[j] > table[j+1] {
				t = table[j+1]
				table[j+1] = table[j]
				table[j] = t
			}
		}
	}
}
