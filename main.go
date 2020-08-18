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
func Shuffle(gs GameState) GameState {
	ret := Clone(gs)
	ret.Deck = deck.NewDeck(deck.MultiDeck(3), deck.ShuffleDec)
	return ret
}

func Deal(gs GameState) GameState {
	ret := Clone(gs)
	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, ret.Deck = ret.Deck[0], ret.Deck[1:]
		ret.Player = append(ret.Player, card)
		card, ret.Deck = ret.Deck[0], ret.Deck[1:]
		ret.Dealer = append(ret.Dealer, card)
	}
	ret.State = StatePlayersTurn
	return ret
}
func Hit(gs GameState) GameState {
	ret := Clone(gs)
	hand := ret.CurrentPlayer()
	var card deck.Card
	card, gs.Deck = gs.Deck[0], gs.Deck[1:]
	*hand = append(*hand, card)
	if hand.Score() > 21 {
		return Stand(ret)
	}
	return ret
}

func Stand(gs GameState) GameState {
	ret := Clone(gs)
	ret.State++
	return ret
}

func FinalHand(gs GameState) GameState {
	ret := Clone(gs)
	fmt.Println("--- Final Hand ---")
	pScore, dScore := ret.Player.Score(), ret.Dealer.Score()
	fmt.Println("Players Cards : ", ret.Player, "\nScore : ", pScore)
	fmt.Println("Dealers Cards : ", ret.Dealer, "\nScore : ", dScore)
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
	ret.Player = nil
	ret.Dealer = nil
	return ret
}
func main() {
	var gs GameState
	gs = Shuffle(gs)

	gs = Deal(gs)
	var input string
	for gs.State == StatePlayersTurn {
		fmt.Println("Players Cards : ", gs.Player)
		fmt.Println("Dealers Cards : ", gs.Dealer.dealerString())
		fmt.Println("(h)it or (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			gs = Hit(gs)
		case "s":
			gs = Stand(gs)
		default:
			fmt.Println("Invalid input")
		}
	}

	for gs.State == StateDealersTurn {
		if (gs.Dealer.Score() <= 16) || (gs.Dealer.Score() == 17 && gs.Dealer.MinScore() != 17) {
			gs = Hit(gs)
		} else {
			gs = Stand(gs)
		}
	}
	gs = FinalHand(gs)
}

type State int8

const (
	StatePlayersTurn State = iota
	StateDealersTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	State  State
	Player Hand
	Dealer Hand
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayersTurn:
		return &gs.Player
	case StateDealersTurn:
		return &gs.Dealer
	default:
		panic("There is no current player turn")

	}
}
func Clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}
