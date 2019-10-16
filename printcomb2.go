package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for t := '0'; t <= '9'; t++ {
					st1 := i*10 + j
					st2 := k*10 + t
					if st1 < st2 && !(i == '9' && j == '8' && k == '9' && t == '9') {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(32)
						z01.PrintRune(k)
						z01.PrintRune(t)
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
	z01.PrintRune('9')
	z01.PrintRune('8')
	z01.PrintRune(32)
	z01.PrintRune('9')
	z01.PrintRune('9')
	z01.PrintRune('\n')
}
