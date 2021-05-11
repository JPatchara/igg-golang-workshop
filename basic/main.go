package main

func main() {
	cards := deck{}
	cards.print(newDeck())
}

func newCard() string {
	return "Ace of Spades"
}

func newDeck() []string {
	generatedCardDeck := []string{}
	cardSuits := []string{"Diamonds", "Hearts", "Spades", "Clubs"}
	cardValues := []string{
		"Ace",
		"Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten",
		"Jack", "Queen", "King",
	}

	for _, suit := range cardSuits {
		for _, cardValue := range cardValues {
			generatedCardDeck = append(generatedCardDeck, suit+" of "+cardValue)
		}
	}

	return generatedCardDeck
}
