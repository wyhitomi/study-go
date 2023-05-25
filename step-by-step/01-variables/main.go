package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"   // shorthand for the above line, go interprets this as a type string
	card = "Five of Diamonds" // reassigning the value of card, reasigning a different type of value will result in an error

	fmt.Println(card)
}
