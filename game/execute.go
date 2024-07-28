package game

import (
	"solitaire/deck"
	"solitaire/utils"
	"strconv"
	"strings"
)

// Moves
// D - [Draw]
// MXX - [Move] [FromPile] [FromPileIndex] [ToPile]
// CX - [Complete] [FromPile]

func (g *game) executeMove() {
	move := strings.ToLower(g.InputBuf.String())

	if len(move) == 0 {
		return
	}

	switch move[0] {
	case byte('d'):
		g.drawCard()
		g.InputBuf.Reset()
	case byte('m'):
		if len(move) == 3 {
			if move[1] == 'o' {
				g.moveCard(string(move[1]), "", string(move[2]))
				g.InputBuf.Reset()
				return
			}
		}

		if len(move) == 4 {
			g.moveCard(string(move[1]), string(move[2]), string(move[3]))
			g.InputBuf.Reset()
			return
		}
	case byte('c'):
		if len(move) != 2 {
			return
		}

		g.completeMove(string(move[1]))
		g.InputBuf.Reset()
		return
	default:
		return
	}
}

func (g *game) drawCard() {
	if len(g.DrawPile.Cards) == 0 {
		if len(g.OpenPile.Cards) == 0 {
			return
		}

		for _, card := range g.OpenPile.Cards {
			card.IsOpen = false
			g.DrawPile.PutCard(card)
		}
	}

	card := g.DrawPile.PopCard()

	card.IsOpen = true

	g.OpenPile.PutCard(card)
}

func (g *game) moveCard(from, fromPileIdx, to string) {
	if !utils.CharInSlice(from, []string{"1", "2", "3", "4", "5", "6", "7", "o", "s", "d", "h", "c"}) {
		return
	}

	if !utils.CharInSlice(to, []string{"1", "2", "3", "4", "5", "6", "7"}) {
		return
	}

	fromDeck := g.getFromDeck(from)

	toInt, err := strconv.Atoi(to)
	if err != nil {
		utils.Log("to int conv error")
	}

	toCard := g.Decks[toInt-1].GetTop()

	if !g.isOppositeColor(fromDeck.GetTop(), toCard) {
		utils.Log("not opposite color")
		return
	}

	if !g.isOrdered(fromDeck.GetTop(), toCard, false) {
		utils.Log("incorrect order")
		return
	}

	var tempPile []*deck.Card

	if fromPileIdx != "" {
		fromPileIdxInt, err := strconv.Atoi(fromPileIdx)
		if err != nil {
			utils.Log("from pile index conv err")
			return
		}

		fromCard := fromDeck.GetIdx(fromPileIdxInt)
		if fromCard == nil {
			utils.Log("from pile idx card nil")
			return
		}

		c := len(fromDeck.Cards) - fromPileIdxInt

		for range c {
			tempPile = append(tempPile, fromDeck.PopCard())
		}
	} else {
		tempPile = []*deck.Card{fromDeck.PopCard()}
	}

	for i := len(tempPile) - 1; i >= 0; i-- {
		g.Decks[toInt-1].PutCard(tempPile[i])
	}
}

func (g *game) isOrdered(fromCard, toCard *deck.Card, isReverse bool) bool {
	fromCardValue := fromCard.GetNumericValue()
	toCardValue := toCard.GetNumericValue()

	if isReverse {
		return fromCardValue == toCardValue+1
	}

	return fromCardValue+1 == toCardValue
}

func (g *game) isOppositeColor(card1, card2 *deck.Card) bool {
	_, ok1 := deck.RedSuite[card1.Suite]
	_, ok2 := deck.RedSuite[card2.Suite]

	if ok1 && ok2 {
		return false
	}

	_, ok1 = deck.BlackSuite[card1.Suite]
	_, ok2 = deck.BlackSuite[card2.Suite]

	if ok1 && ok2 {
		return false
	}

	return true
}

func (g *game) completeMove(from string) {
	if !utils.CharInSlice(from, []string{"1", "2", "3", "4", "5", "6", "7", "o"}) {
		return
	}

	fromDeck := g.getFromDeck(from)

	switch fromDeck.GetTop().Suite {
	case deck.Clubs:
		g.moveCompletedCard(&g.ClubsPile, fromDeck)
		return
	case deck.Hearts:
		g.moveCompletedCard(&g.HeartsPile, fromDeck)
		return
	case deck.Diamonds:
		g.moveCompletedCard(&g.DiamondsPile, fromDeck)
		return
	case deck.Spades:
		g.moveCompletedCard(&g.SpadesPile, fromDeck)
		return
	}
}

func (g *game) moveCompletedCard(complete *deck.Deck, from *deck.Deck) {
	if !g.isOrdered(from.GetTop(), complete.GetTop(), true) {
		utils.Log("not ordered")
		return
	}

	popCard := from.PopCard()
	complete.PutCard(popCard)
}

func (g *game) getFromDeck(from string) *deck.Deck {
	switch from {
	case "1", "2", "3", "4", "5", "6", "7":
		fromInt, err := strconv.Atoi(from)
		if err != nil {
			utils.Log("from int conv error")
		}

		return g.returnNonEmptyPile(&g.Decks[fromInt-1])
	case "o":
		return g.returnNonEmptyPile(&g.OpenPile)
	case "s":
		return g.returnNonEmptyPile(&g.SpadesPile)
	case "d":
		return g.returnNonEmptyPile(&g.DiamondsPile)
	case "h":
		return g.returnNonEmptyPile(&g.HeartsPile)
	case "c":
		return g.returnNonEmptyPile(&g.ClubsPile)
	default:
		utils.Log("cannot find from pile")
		return nil
	}
}

func (g *game) returnNonEmptyPile(pile *deck.Deck) *deck.Deck {
	if pile.Cards == nil || len(pile.Cards) == 0 {
		utils.Log("pile empty")
		return nil
	}

	return pile
}
