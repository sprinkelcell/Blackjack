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

func (hand Hand) Score() int {
	minScore := hand.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, val := range hand {
		if val.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}
func (hand Hand) MinScore() int {
	score := 0
	for _, val := range hand {
		score += min(int(val.Rank), 10)
	}
	return score
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
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

	for (dealer.Score() <= 16) || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, cards = cards[0], cards[1:]
		dealer = append(dealer, card)
	}
	fmt.Println("--- Final Hand ---")
	pScore, dScore := player.Score(), dealer.Score()
	fmt.Println("Players Cards : ", player, "\nScore : ", pScore)
	fmt.Println("Dealers Cards : ", dealer, "\nScore : ", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted!")
	case dScore > 21:
		fmt.Println("Dealer busted!")
	case pScore > dScore:
		fmt.Println("You win")
	case pScore < dScore:
		fmt.Println("You lose")
	case pScore == dScore:
		fmt.Println("Draw")

	}
}
