package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n != -9223372036854775808 {
		tt := 1
		pp := 1
		bb := 0
		for i := 0; i < 18; i++ {
			if n < 0 {
				n = -1 * n
				z01.PrintRune('-')
			}
			tt = 10 * tt
			pp = i + 1
			if n < tt {
				break
			}
		}

		if pp == 18 {
			bb = 1
			pp = 19
		}
		for i := 1; i <= pp; i++ {
			if bb == 1 {
				sym := n / tt
				n = n - sym*tt
				if sym == 0 {
					z01.PrintRune('0')
				}
				if sym == 1 {
					z01.PrintRune('1')
				}
				if sym == 2 {
					z01.PrintRune('2')
				}
				if sym == 3 {
					z01.PrintRune('3')
				}
				if sym == 4 {
					z01.PrintRune('4')
				}
				if sym == 5 {
					z01.PrintRune('5')
				}
				if sym == 6 {
					z01.PrintRune('6')
				}
				if sym == 7 {
					z01.PrintRune('7')
				}
				if sym == 8 {
					z01.PrintRune('8')
				}
				if sym == 9 {
					z01.PrintRune('9')
				}
				bb = 0
			} else {
				tt = tt / 10
				sym := n / tt
				n = n - sym*tt
				if sym == 0 {
					z01.PrintRune('0')
				}
				if sym == 1 {
					z01.PrintRune('1')
				}
				if sym == 2 {
					z01.PrintRune('2')
				}
				if sym == 3 {
					z01.PrintRune('3')
				}
				if sym == 4 {
					z01.PrintRune('4')
				}
				if sym == 5 {
					z01.PrintRune('5')
				}
				if sym == 6 {
					z01.PrintRune('6')
				}
				if sym == 7 {
					z01.PrintRune('7')
				}
				if sym == 8 {
					z01.PrintRune('8')
				}
				if sym == 9 {
					z01.PrintRune('9')
				}
			}
		}
	} else {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
	}
}
