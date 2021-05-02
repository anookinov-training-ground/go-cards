package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards := newDeckFromFile("my_cards")
	// cards.print()
	// var card string = "Ace of Spades"
	// var card = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()
	// fmt.Println(card)
	// cards := []string{"Ace of Diamonds", newCard()}
	// fmt.Println(cards)
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	// cards := newDeck()
	// cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
}

// func newCard() string {
// 	return "Five of Diamonds"
// }