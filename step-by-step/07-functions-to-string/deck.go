package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Returns a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// Prints the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)

	}
}

// Returns a dealed hand
func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

// converts a deck to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	// return []string(d)
}
