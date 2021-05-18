package main

import "testing"

func testNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of deck is Ace of Spades, but got %v", d[0])
	}

	if d[0] != "Four of Clubs" {
		t.Errorf("Expected last card of deck is Four of Clubs, but got %v", d[0])
	}
}
