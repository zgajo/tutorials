package main

func main() {
	// Two ways to define variables,
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // only to define new variable

	// array vs slice
	// Array - fixed length list of things
	// Slice - array that cang grow and shrink
	// slice
	cards := newDeck()

	cards.shuffle()
	cards.print()
	// hand, remainingDeck := cards.deal(5)

	// hand.print()
	// remainingDeck.print()

	// fmt.Println(hand.toString())
	// hand.savetoFile("./01_cards_project/cards.txt")

	// cards2 := newDeckFromFile("./01_cards_project/cards.txt")
	// cards2.print()
	// fmt.Println([]byte("Hello"))
}
