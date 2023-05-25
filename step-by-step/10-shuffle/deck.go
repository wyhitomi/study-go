package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
}

// save to file
// for learning purposes we are using the ioutil package, however it was deprecated since go 1.16
// for more information take a look at https://pkg.go.dev/io/ioutil
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// read from file
// for learning purposes we are using the ioutil package, however it was deprecated since go 1.16
// for more information take a look at https://pkg.go.dev/io/ioutil
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// shuffle a deck using Unix time as seed to emulate a true random number generator
func (d deck) shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
