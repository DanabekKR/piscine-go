package piscine

func Sqrt(nb int) int {
	for i := 0; i <= nb; i++ {
		if i*i == nb {
			return i
		}
		if i*i > nb {
			return 0
		}
	}
	return 0
}
