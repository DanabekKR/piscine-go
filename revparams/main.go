package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	count := 0
	for range arguments {
		count++
	}
	for i := count - 1; i > 0; i-- {
		pn := arguments[i]
		if i > 0 {
			for _, a := range pn {
				z01.PrintRune(a)
			}
		}
		z01.PrintRune('\n')
	}
}
