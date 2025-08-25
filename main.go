package main

import (
	"fmt"

	"github.com/RodrigoCelso/gophercises-09/deck"
)

func main() {
	deckSort := deck.New(deck.WithJoker(2), deck.WithSort())
	deckSortFunc := deck.New(
		deck.WithJoker(1),
		deck.WithSortFunc(func(cards []deck.Card) func(i int, j int) bool {
			return func(i, j int) bool {
				return cards[i].Value < cards[j].Value
			}
		}))
	deckShuffle := deck.New(deck.WithShuffle())

	fmt.Println("Deck Sort:", deckSort)
	fmt.Println("Deck Sort Func:", deckSortFunc)
	fmt.Println("Deck Shuffle:", deckShuffle)

}
