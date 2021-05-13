package main

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 3)

	// fmt.Println("Hand cards_____:")
	// hand.print()

	// fmt.Println("Remaining cards_____:")
	// remainingCards.print()

	// if err := cards.SaveToFile("deckInfo.txt"); err != nil {
	// 	fmt.Println("Error : ", err)
	// }

	cards := newDeckFromFile("deckInfo.txt")
	cards.shuffle()
	cards.print()
}
