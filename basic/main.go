package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingCard := deal(cards, 3)

	fmt.Println("Hand cards_____:")
	hand.print()

	fmt.Println("Remaining cards_____:")
	remainingCard.print()
}
