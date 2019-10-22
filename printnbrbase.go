package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	n := StrLen(base)
	isNV := false
	if n < 1 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		isNV = true
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if base[i] == base[j] && !isNV {
				z01.PrintRune('N')
				z01.PrintRune('V')
				isNV = true
				break
			}
		}
	}
	if !isNV {
		var list []int
		isNeg := false
		if nbr < 0 {
			nbr = nbr * (-1)
			isNeg = true
		}
		for i := 0; i < 100; i++ {
			if nbr < n {
				list = append(list, nbr)
				break
			}
			list = append(list, nbr%n)
			nbr = nbr / n
		}
		count := 0
		for range list {
			count++

		}
		if isNeg {
			z01.PrintRune('-')
		}
		for i := count - 1; i >= 0; i-- {
			tmp := list[i]
			z01.PrintRune(rune(base[tmp]))
		}
	}
}
