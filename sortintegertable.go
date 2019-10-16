package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if table[i] > table[j] {
				Swap(&table[i], &table[j])
			}
		}
	}
}
