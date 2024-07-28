package game

import (
	"fmt"
	"solitaire/deck"
	"solitaire/screen"
	"strings"
)

type Game interface {
	Start()
}

type game struct {
	IsRunning    bool
	numMoves     int
	score        int
	InputBuf     *strings.Builder
	screen       screen.Screen
	Decks        [7]deck.Deck
	DrawPile     deck.Deck
	OpenPile     deck.Deck
	DiamondsPile deck.Deck
	SpadesPile   deck.Deck
	HeartsPile   deck.Deck
	ClubsPile    deck.Deck
}

func NewGame() Game {
	newDeck := deck.NewDeck()

	fmt.Printf("deck: %+v\n", newDeck)

	decks := [7]deck.Deck{}

	for i := 1; i <= 7; i++ {
		cards := []*deck.Card{}
		for x := range i {
			fmt.Printf("i: %d, x: %d\n", i, x)
			c := newDeck.PopCard()
			if x == i-1 {
				c.IsOpen = true
			}
			cards = append(cards, c)
		}

		decks[i-1] = deck.NewFromCards(cards)
	}

	drawPile := newDeck

	return &game{
		IsRunning:    false,
		numMoves:     0,
		score:        0,
		InputBuf:     &strings.Builder{},
		screen:       screen.NewScreen(),
		Decks:        decks,
		DrawPile:     drawPile,
		OpenPile:     deck.NewFromCards([]*deck.Card{}),
		DiamondsPile: deck.NewFromCards([]*deck.Card{}),
		ClubsPile:    deck.NewFromCards([]*deck.Card{}),
		HeartsPile:   deck.NewFromCards([]*deck.Card{}),
		SpadesPile:   deck.NewFromCards([]*deck.Card{}),
	}
}
