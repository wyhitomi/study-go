package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	// array := [2]string // arrays have fixed length, in this case you will get a 2 position array

	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {

	return "Five of Diamonds"
}
