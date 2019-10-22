package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	n := 0
	for range base {
		n++
	}
	isNV := false
	if n < 2 {
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
		var list [100]int
		for i := range list {
			list[i] = -1
		}
		isNeg := false
		if nbr < 0 {
			nbr = nbr * (-1)
			isNeg = true
		}
		app := 0
		for i := 0; i < 100; i++ {
			if nbr < n {
				list[app] = nbr
				break
			}
			list[app] = nbr % n
			nbr = nbr / n
			app++
		}
		count := 0
		for _, a := range list {
			if !(a == -1) {
				count++
			}
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
