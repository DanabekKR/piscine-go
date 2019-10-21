package piscine

func IsPrime(nb int) bool {
	sqVal := 0
	if nb < 0 {
		nb = nb * (-1)
	}
	for i := 0; i <= nb; i++ {
		if i*i == nb {
			return false
		}
		if i*i > nb {
			sqVal = i - 1
			break
		}
	}
	for i := 2; i <= sqVal; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
