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
	deckFilter := deck.New(
		deck.WithJoker(5),
		deck.WithFilter(func(card deck.Card) bool {
			return card.Value == deck.Jack || card.Suit == deck.Joker
		}))
	deckMultiple := deck.New(deck.WithMultipleDeck(*deckSort, *deckSortFunc, *deckShuffle, *deckFilter), deck.WithSort())

	fmt.Println("Deck Sort:", deckSort)
	fmt.Println("Deck Sort Func:", deckSortFunc)
	fmt.Println("Deck Shuffle:", deckShuffle)
	fmt.Println("Deck Filter:", deckFilter)
	fmt.Println("Deck Multiple:", deckMultiple)
}
