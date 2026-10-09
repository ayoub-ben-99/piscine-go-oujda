package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	player := 1
	card := 0

	for _, value := range deck {
		if card == 0 {
			fmt.Printf("Player %d: ", player)
		}

		fmt.Printf("%d", value)

		card++

		if card == 3 {
			z01.PrintRune('\n')
			player++
			card = 0
		} else {
			fmt.Printf(", ")
		}
	}
}
