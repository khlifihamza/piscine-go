package piscine

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if len(slice[i]) < len(slice[j]) {
				t := slice[j]
				slice[j] = slice[i]
				slice[i] = t
			}
		}
	}
	return slice
}
