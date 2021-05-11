package main

import "fmt"

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	generatedCardDeck := deck{}
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

func deal(cards deck, handedSize int) (deck, deck) {
	return cards[:handedSize], cards[handedSize:]
}
