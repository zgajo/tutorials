package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		// %v - passed variable as argument (len(d))
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of spades, but got %v", d[0])
	}

}

func TestSaveToDeckAndTestNewDeckFromFile(t *testing.T) {
	testingFile := "_desting.txt"
	os.Remove(testingFile)

	d := newDeck()
	d.savetoFile(testingFile)

	loadedDeck := newDeckFromFile(testingFile)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	os.Remove(testingFile)
}
