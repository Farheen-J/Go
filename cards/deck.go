package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
Create a new type of 'deck'
which is a slice of strings.
Type 'deck' extends all the behavior of slice of strings
*/
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"} //Enough to test

	for _, suit := range cardSuits { //Replace index with _
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

/*
Select range of cards from deck to return hand of deck and remaining decks
Not a receiver function because func (d deck) deal...
*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/*
Convert deck to byte_slice[ASCII ELEMENTS] in the process
deck -> []string -> single string
Receiver
*/
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

/*
Save the deck to file.
Convert deck to byte_slice[ASCII ELEMENTS] in the process
single string -> []byte
Receiver
*/
func (d deck) savetoFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 777)
}

/*
Read the deck to file.
[]byte -> string -> []string, deck
*/
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

/*
Shuffle
*/
func (d deck) shuffle() {

	//pass seed to random value generator
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		//Swapping
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

/*
Receiver function: sets up methods on variables.
d = actual copy of the deck (object).

	By convention, this has to be one/two lettered abbreviation which matches the type but not mandatory

deck = every variable of type deck (class)
We only add a receiver if we want to call object.func()
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
