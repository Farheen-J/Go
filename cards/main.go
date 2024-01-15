package main

func main() {
	cards := newDeck()

	// hand, remainingCards := deal(cards, 5) //No import needed. If we are passing object, no need to use object.func() syntax
	// fmt.Println("Hand Deck: ")
	// hand.print()
	// fmt.Println("Remaining Deck: ")
	// remainingCards.print()
	//cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}
