package main

// type deck []string

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	cards := newDeck()

	// for index, card := range cards {
	// 	fmt.Println(index, card)
	// }
	cards.print() // using the print() function declared on deck.go
}
