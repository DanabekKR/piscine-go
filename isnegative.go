package piscine

import "github.com/01-edu/z01"

func IsNegative(n int) {
	if n < 0 {
		z01.PrintRune('T')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('F')
		z01.PrintRune('\n')
	}
}
