package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	tt := 1
	pp := 1
	for i := 0; i < 100; i++ {
		if n < 0 {
			n = -1 * n
			z01.PrintRune('-')
		}
		tt = 10 * tt
		if (n / tt) < 10 {
			pp = i + 1
			sym := n / tt
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
			break
		}
	}
	for i := 1; i <= pp; i++ {
		if tt > 10 {
			tt = tt / 10
		}
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
