package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var commaSeparator string = ", "

type deck []string

func newDeck() deck {
	cards := deck{}

	// cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	var cardSuits = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cs := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+cs)
		}
	}

	return cards
}

func (d deck) print() {
	// range = iteare over every elemtent in cards
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), commaSeparator)
}

func (d deck) savetoFile(fileName string) error {
	var byteSlice = []byte(d.toString())
	return ioutil.WriteFile(fileName, byteSlice, 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cardsString := string(byteSlice)
	cardsSlice := strings.Split(cardsString, commaSeparator)
	deck := deck(cardsSlice)

	return deck

}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {

		randPosition := r.Intn(len(d) - 1)

		d[i], d[randPosition] = d[randPosition], d[i]

	}

	return d
}
