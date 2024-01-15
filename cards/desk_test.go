package main

import (
	"os"
	"testing"
)

/*
function name starts with uppercase:
t:= test handler
We don't get number of test cases
*/
func TestNewDeck(t *testing.T) {
	d := newDeck()

	//test1
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	//test2
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	//test3
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	//Remove reminents of previously failed tests
	os.Remove("_decktesting")

	deck := newDeck()
	deck.savetoFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
