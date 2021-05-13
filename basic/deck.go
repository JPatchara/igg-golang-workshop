package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string
type Rand struct{}

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

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0777)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error!", err)
	}
	return strings.Split(string(bs), ",")
}

func (d deck) shuffle() {
	rSource := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		newPosition := rSource.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
