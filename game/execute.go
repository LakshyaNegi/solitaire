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

func (g *Game) ExecuteMove() {
	move := strings.ToLower(g.InputBuf.String())

	if len(move) == 0 {
		return
	}

	switch move[0] {
	case byte('d'):
		g.DrawCard()
		g.InputBuf.Reset()
	case byte('m'):
		if len(move) == 3 {
			if move[1] == 'o' {
				g.MoveCard(string(move[1]), "", string(move[2]))
				g.InputBuf.Reset()
				return
			}
		}

		if len(move) == 4 {
			g.MoveCard(string(move[1]), string(move[2]), string(move[3]))
			g.InputBuf.Reset()
			return
		}
	case byte('c'):
		if len(move) != 2 {
			return
		}

		g.CompleteMove(string(move[1]))
		g.InputBuf.Reset()
		return
	default:
		return
	}
}

func (g *Game) DrawCard() {
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

func (g *Game) MoveCard(from, fromPileIdx, to string) {
	var fromDeck *deck.Deck
	switch from {
	case "1", "2", "3", "4", "5", "6", "7":
		fromInt, err := strconv.Atoi(from)
		if err != nil {
			utils.Log("from int conv error")
		}

		if g.Decks[fromInt-1].Cards == nil || len(g.Decks[fromInt-1].Cards) == 0 {
			utils.Log("from pile is empty")
			return
		}

		fromDeck = &g.Decks[fromInt-1]
	case "o":
		if g.OpenPile.Cards == nil || len(g.OpenPile.Cards) == 0 {
			utils.Log("open pile empty")
			return
		}

		fromDeck = &g.OpenPile
	case "s":
		if g.SpadesPile.Cards == nil || len(g.SpadesPile.Cards) == 0 {
			utils.Log("spades pile empty")
			return
		}

		fromDeck = &g.SpadesPile
	case "d":
		if g.DiamondsPile.Cards == nil || len(g.DiamondsPile.Cards) == 0 {
			utils.Log("spades pile empty")
			return
		}

		fromDeck = &g.DiamondsPile
	case "h":
		if g.HeartsPile.Cards == nil || len(g.HeartsPile.Cards) == 0 {
			utils.Log("spades pile empty")
			return
		}

		fromDeck = &g.HeartsPile
	case "c":
		if g.ClubsPile.Cards == nil || len(g.ClubsPile.Cards) == 0 {
			utils.Log("spades pile empty")
			return
		}

		fromDeck = &g.ClubsPile
	default:
		utils.Log("cannot find from pile")
		return
	}

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

	switch to {
	case "1", "2", "3", "4", "5", "6", "7":

		toInt, err := strconv.Atoi(to)
		if err != nil {
			utils.Log("to int conv error")
		}

		toCard := g.Decks[toInt-1].GetTop()

		if !g.IsOppositeColor(fromCard, toCard) {
			utils.Log("not opposite color")
			return
		}

		if !g.IsOrdered(fromCard, toCard) {
			utils.Log("incorrect order")
			return
		}

		c := len(fromDeck.Cards) - fromPileIdxInt
		tempPile := []*deck.Card{}

		for range c {
			tempPile = append(tempPile, fromDeck.PopCard())
		}

		// popCard := fromDeck.PopCard()
		for i := len(tempPile) - 1; i >= 0; i-- {
			g.Decks[toInt-1].PutCard(tempPile[i])
		}

	default:
		utils.Log("cannot find to pile")
		return
	}

}

func (g *Game) IsOrdered(fromCard, toCard *deck.Card) bool {
	switch fromCard.Value {
	case deck.King:
		if toCard != nil {
			return false
		}

	case deck.Queen:
		if toCard.Value != deck.King {
			return false
		}

	case deck.Jack:
		if toCard.Value != deck.Queen {
			return false
		}

	case deck.Ten:
		if toCard.Value != deck.Jack {
			return false
		}

	case deck.Nine:
		if toCard.Value != deck.Ten {
			return false
		}

	case deck.Eight:
		if toCard.Value != deck.Nine {
			return false
		}

	case deck.Seven:
		if toCard.Value != deck.Eight {
			return false
		}
	case deck.Six:
		if toCard.Value != deck.Seven {
			return false
		}

	case deck.Five:
		if toCard.Value != deck.Six {
			return false
		}

	case deck.Four:
		if toCard.Value != deck.Five {
			return false
		}

	case deck.Three:
		if toCard.Value != deck.Four {
			return false
		}

	case deck.Two:
		if toCard.Value != deck.Three {
			return false
		}

	case deck.Ace:
		if toCard.Value != deck.Two {
			return false
		}
	}

	return true

}

func (g *Game) IsOppositeColor(card1, card2 *deck.Card) bool {
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

func (g *Game) CompleteMove(from string) {
	var fromDeck *deck.Deck
	switch from {
	case "1", "2", "3", "4", "5", "6", "7":
		fromInt, err := strconv.Atoi(from)
		if err != nil {
			utils.Log("from int conv error")
		}

		if g.Decks[fromInt-1].Cards == nil || len(g.Decks[fromInt-1].Cards) == 0 {
			utils.Log("from pile is empty")
			return
		}

		fromDeck = &g.Decks[fromInt-1]
	case "o":
		if g.OpenPile.Cards == nil || len(g.OpenPile.Cards) == 0 {
			utils.Log("open pile empty")
			return
		}

		fromDeck = &g.OpenPile
	default:
		utils.Log("cannot find from pile")
		return
	}

	fromCard := fromDeck.GetTop()

	switch fromCard.Suite {
	case deck.Clubs:
		topCard := g.ClubsPile.GetTop()

		if topCard.GetNumericValue()+1 != fromCard.GetNumericValue() {
			utils.Log("cannot complete the card, not ordered")
			return
		}

		popCard := fromDeck.PopCard()
		g.ClubsPile.PutCard(popCard)
		return
	case deck.Hearts:
		topCard := g.HeartsPile.GetTop()

		if topCard.GetNumericValue()+1 != fromCard.GetNumericValue() {
			utils.Log("cannot complete the card, not ordered")
			return
		}

		popCard := fromDeck.PopCard()
		g.HeartsPile.PutCard(popCard)
		return
	case deck.Diamonds:
		topCard := g.DiamondsPile.GetTop()

		if topCard.GetNumericValue()+1 != fromCard.GetNumericValue() {
			utils.Log("cannot complete the card, not ordered")
			return
		}

		popCard := fromDeck.PopCard()
		g.DiamondsPile.PutCard(popCard)
		return
	case deck.Spades:
		topCard := g.SpadesPile.GetTop()

		if topCard.GetNumericValue()+1 != fromCard.GetNumericValue() {
			utils.Log("cannot complete the card, not ordered")
			return
		}

		popCard := fromDeck.PopCard()
		g.SpadesPile.PutCard(popCard)
		return
	}
}
