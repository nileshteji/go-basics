package main

import (
	"fmt"
	"strings"
)

type deck []string

func main() {
	cards := newDeck()
	fmt.Println(cards.toByte())

}

func (d deck) toByte() []byte {
	return []byte(strings.Join(d, " "))
}

func (d deck) hand(variable int) (deck, deck) {
	return d[:variable], d[variable:]
}


func newDeck() deck {

	cardSuites := deck{"Spades", "Hearts", "Diamonds", "Leaves"}
	cardValues := deck{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jacks", "Queen", "King"}
	cards := deck{}

	for _, cardS := range cardSuites {
		for _, cardV := range cardValues {
			cards = append(cards, cardS+" of "+cardV)
		}
	}

	return cards

}
