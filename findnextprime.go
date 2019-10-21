package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	for i := nb; i <= 2*nb; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 2
}
