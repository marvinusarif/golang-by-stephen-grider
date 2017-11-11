package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of deck
//which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Cubs"}
	cardValues := []string{"Ace", "Jack", "Queen", "King", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//(d deck) as receiver
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
	//or
	//return strings.Split(string(bs), ",")
}

func (d deck) shuffle() {
	t := time.Now().UnixNano()
	source := rand.NewSource(t)
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		//swapping the element
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
