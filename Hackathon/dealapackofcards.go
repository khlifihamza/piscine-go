package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	l := 0
	for range deck {
		l++
	}
	j := 0
	for i := 1; i <= l/3; i++ {
		fmt.Printf("Player %v: %v, %v, %v", i, deck[j], deck[j+1], deck[j+2])
		z01.PrintRune('\n')
		j += l / 4
	}
}
