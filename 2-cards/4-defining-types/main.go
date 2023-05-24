package main

import "fmt"

// type deck []string

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	cards := deck
	cards = append(cards, "Six of Spades")

	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {

	return "Five of Diamonds"
}
