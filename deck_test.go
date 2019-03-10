package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length to be 52, received %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades' as first card in deck, received %v", d[0])
	}

	if d[len(d)-1] != "Queen of Clubs" {
		t.Errorf("Expected 'Queen of Clubs' as first card in deck, received %v", d[0])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length to be 52, received %v", len(loadedDeck))
	}

	if loadedDeck[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades' as first card in deck, received %v", loadedDeck[0])
	}

	if loadedDeck[len(loadedDeck)-1] != "Queen of Clubs" {
		t.Errorf("Expected 'Queen of Clubs' as first card in deck, received %v", loadedDeck[0])
	}

	os.Remove("_deckTesting")
}
