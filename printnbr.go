package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	tt := 1
	pp := 1
	for i := 0; i < 15; i++ {
		if n < 0 {
			n = -1 * n
			z01.PrintRune('-')
		}
		tt = 10 * tt
		if n < tt {
			pp = i + 1
			break
		}
	}
	for i := 1; i <= pp; i++ {
		tt = tt / 10
		sym := n / tt
		n = n - sym*tt
		for j := '0'; j <= '9'; j++ {
			if sym+48 == int(j) {
				z01.PrintRune(j)
			}
		}
	}
}
