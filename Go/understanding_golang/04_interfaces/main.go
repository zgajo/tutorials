package main

import "fmt"

type bot interface {
	// If there is a type with a function called getGreeting and returns string
	// then you're member of type interface bot
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello"
}

func (sb spanishBot) getGreeting() string {
	return "Hola"
}
