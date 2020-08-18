package main

import (
	"fmt"
	"strings"

	"github.com/sprinkelcell/deck"
)

type Hand []deck.Card

func (hand Hand) String() string {
	str := make([]string, len(hand))
	for index := range hand {
		str[index] = hand[index].String()
	}
	return strings.Join(str, ", ")
}

func (hand Hand) dealerString() string {
	return hand[0].String() + ", **HIDDEN**"
}
func main() {
	cards := deck.NewDeck(deck.MultiDeck(3), deck.ShuffleDec)
	var card deck.Card
	var player, dealer Hand
	for i := 0; i < 2; i++ {

		for _, hand := range []*Hand{&player, &dealer} {
			card, cards = cards[0], cards[1:]
			*hand = append(*hand, card)
		}
	}
	var input string
	for input != "s" {
		fmt.Println("Players Cards : ", player)
		fmt.Println("Dealers Cards : ", dealer.dealerString())
		fmt.Println("(h)it or (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, cards = cards[0], cards[1:]
			player = append(player, card)
		case "s":
			continue
		default:
			fmt.Println("Invalid input")
		}
	}
	fmt.Println("--- Final Hand ---")
	fmt.Println("Players Cards : ", player)
	fmt.Println("Dealers Cards : ", dealer)
}
