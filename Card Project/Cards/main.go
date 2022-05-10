package main

func main() {
	// var card string = "Ace of Spades"
	// := for initialisation
	// // multiple returns of one function
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
	//card = "Five of Diamonds"
	//cards.print()
	//cards := newDeckFromFile("my_cards")
	//cards.print()
	//fmt.Println(cards.toString())
	//greeting := "Hi there!"
	//fmt.Println([]byte(greeting)) // byte slice of string
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
