package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for i := 1; i < len(arguments); i++ {
		pn := arguments[i]
		if i > 0 {
			for _, a := range pn {
				z01.PrintRune(a)
			}
		}
		z01.PrintRune('\n')
	}
}
