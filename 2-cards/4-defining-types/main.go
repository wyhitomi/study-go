package main

// type deck []string

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	cards := deck{"Ace of Diamonds", "Two of Diamonds"} // using the type deck declared on deck.go
	cards = append(cards, "Six of Spades")

	// for index, card := range cards {
	// 	fmt.Println(index, card)
	// }
	cards.print() // using the print() function declared on deck.go
}

func newCard() string {

	return "Five of Diamonds"
}
