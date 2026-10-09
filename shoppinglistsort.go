package piscine

func ShoppingListSort(slice []string) []string {
	for i := range slice {
		for j := range slice {
			if len(slice[i]) < len(slice[j]) {
				slice[j], slice[i] = slice[i], slice[j]
			}
		}
	}
	return slice
}
