package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, r := range a {
		for _, v := range r {
			z01.PrintRune(v)
		}

		z01.PrintRune('\n')
	}
}
