package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	tmp := 10
	var list []int
	countNum := 0
	t := n
	for i := 0; i < 10; i++ {
		if t < 10 {
			countNum++
			break
		}
		t = t / tmp
		countNum++
	}
	for i := 0; i < countNum; i++ {
		nbr := n % 10
		list = append(list, nbr)
		n = n / 10
	}
	SortIntegerTable(list)
	for i := 0; i < countNum; i++ {
		z01.PrintRune(rune(48 + list[i]))
	}
}
