package main

import (
	"fmt"

	"github.com/RodrigoCelso/gophercises-09/deck"
)

func main() {
	// New Deck Example
	deckNew := deck.New()
	// Deck with Default Sort Example
	deckSort := deck.New(deck.WithSort())
	// Deck with Sort Example
	deckSortFunc := deck.New(
		deck.WithJoker(1),
		deck.WithSortFunc(func(cards []deck.Card) func(i int, j int) bool {
			return func(i, j int) bool {
				return cards[i].Value < cards[j].Value
			}
		}))
	// Deck with Joker Example
	deckJoker := deck.New(deck.WithJoker(3))
	// Deck with Shuffle Example
	deckShuffle := deck.New(deck.WithShuffle())
	// Deck with Filter Example
	deckFilter := deck.New(
		deck.WithJoker(5),
		deck.WithFilter(func(card deck.Card) bool {
			return card.Value == deck.Jack || card.Suit == deck.Joker
		}))
	// Deck with Multiple Decks Example
	deckMultiple := deck.New(deck.WithMultipleDeck(*deckSort, *deckSortFunc, *deckShuffle, *deckFilter), deck.WithSort())

	fmt.Println("New Deck:", deckNew)
	fmt.Println("================================================================================================================")
	fmt.Println("Deck With Sort:", deckSort)
	fmt.Println("================================================================================================================")
	fmt.Println("Deck With Sort Func:", deckSortFunc)
	fmt.Println("================================================================================================================")
	fmt.Println("Deck With Joker", deckJoker)
	fmt.Println("================================================================================================================")
	fmt.Println("Deck With Shuffle:", deckShuffle)
	fmt.Println("================================================================================================================")
	fmt.Println("Deck With Filter:", deckFilter)
	fmt.Println("================================================================================================================")
	fmt.Println("Deck With Multiple:", deckMultiple)
}
