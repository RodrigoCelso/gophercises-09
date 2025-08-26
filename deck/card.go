//go:generate stringer -type=CardSuit,CardValue
package deck

import (
	"math/rand"
	"sort"
)

type CardSuit int

const (
	Spade CardSuit = iota
	Diamond
	Club
	Heart
	suitSize
	Joker
)

type CardValue int

const (
	_ CardValue = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	valueSize
)

type Card struct {
	Suit       CardSuit
	Value      CardValue
	TrumpValue int
}

func newSuitDeck(suit CardSuit) []Card {
	var suitDeck []Card
	for value := range valueSize - 1 {
		card := Card{Suit: suit, Value: CardValue(value + 1)}
		suitDeck = append(suitDeck, card)
	}
	return suitDeck
}

func New(opts ...Option) *[]Card {
	var deck []Card
	for suit := range suitSize {
		deck = append(deck, newSuitDeck(CardSuit(suit))...)
	}
	for _, o := range opts {
		deck = o(deck)
	}
	return &deck
}

type Option func([]Card) []Card

func WithJoker(quantity int) Option {
	return func(c []Card) []Card {
		for idx := 0; idx < quantity; idx++ {
			card := Card{Suit: Joker, Value: CardValue(0), TrumpValue: 55}
			c = append(c, card)
		}
		return c
	}
}

func WithSort() Option {
	return func(c []Card) []Card {
		sort.Slice(c, func(i, j int) bool {
			return c[i].Suit < c[j].Suit
		})
		return c
	}
}

func WithSortFunc(less func(cards []Card) func(i, j int) bool) Option {
	return func(c []Card) []Card {
		sort.Slice(c, less(c))
		return c
	}
}

func WithShuffle() Option {
	return func(c []Card) []Card {
		rand.Shuffle(len(c), func(i, j int) {
			c[i], c[j] = c[j], c[i]
		})
		return c
	}
}

func WithFilter(cardFilter func(card Card) bool) Option {
	return func(cards []Card) []Card {
		var filteredDeck []Card
		for _, card := range cards {
			if !cardFilter(card) {
				filteredDeck = append(filteredDeck, card)
			}
		}
		return filteredDeck
	}
}

func WithMultipleDeck(decks ...[]Card) Option {
	return func(c []Card) []Card {
		var multipleDeck []Card
		for _, deck := range decks {
			multipleDeck = append(multipleDeck, deck...)
		}
		return multipleDeck
	}
}
