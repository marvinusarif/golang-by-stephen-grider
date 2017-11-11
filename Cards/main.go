package main

func main() {
	//var card string = "Ace of spades"
	//card := "Ace of spades"
	//newCard := myNewCard()
	//cards := []string{myNewCard(), myNewCard()}
	//cards := deck{"Ace of Diamonds", myNewCard()}
	//cards = append(cards, "Six of Spades")
	//to update the card variable
	//card = "5 of Diamonds"
	cards := newDeckFromFile("decks")
	//fmt.Println(cards.toString())
	cards.shuffle()
	cards.print()
	cards.saveToFile("decks")
	//hand, remainingDeck := deal(cards, 5)
	//cards.print()
	//hand.print()
	//remainingDeck.print()
}

/*
func myNewCard() string {
	return "Five of Diamonds"
}*/
