package game

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"golang.org/x/term"
)

func (g *Game) Draw() {
	if len(g.ClubsPile.Cards) == 13 &&
		len(g.HeartsPile.Cards) == 13 &&
		len(g.DiamondsPile.Cards) == 13 &&
		len(g.SpadesPile.Cards) == 13 {
		g.IsRunning = false
	}

	g.screen.Clear()

	x, _, err := term.GetSize(0)
	if err != nil {
		log.Fatal(err)
	}

	spadeTop := g.SpadesPile.GetTop()
	g.screen.DrawAt("Spades(S): "+spadeTop.Ansi(), 0, 1)

	diamondsTop := g.DiamondsPile.GetTop()
	g.screen.DrawAt("Diamonds(D): "+diamondsTop.Ansi(), 0, 2)

	clubsTop := g.ClubsPile.GetTop()
	g.screen.DrawAt("Clubs(C): "+clubsTop.Ansi(), 0, 3)

	heartsTop := g.HeartsPile.GetTop()
	g.screen.DrawAt("Hearts(H): "+heartsTop.Ansi(), 0, 4)

	openPileTop := g.OpenPile.GetTop()
	g.screen.DrawAt("Open(O): "+openPileTop.Ansi(), x/2, 1)

	if len(g.DrawPile.Cards) != 0 {
		g.screen.DrawAt("Draw(D): "+"##", x/2, 2)
	} else {
		g.screen.DrawAt("Draw: "+"00", x/2, 2)
	}

	g.screen.DrawAt("Moves: "+strconv.Itoa(g.numMoves), x/2, 3)
	g.screen.DrawAt("Score: "+strconv.Itoa(g.score), x/2, 4)

	for i := range x {
		g.screen.DrawAt("-", i, 6)
	}

	width := x / 6
	i := 0
	c := 1

	for i < x {
		g.screen.DrawAt("--"+strconv.Itoa(c)+"-", i, 7)

		l := len(g.Decks[c-1].Cards)
		for idx, card := range g.Decks[c-1].Cards {
			if idx == l-1 {
				card.IsOpen = true
			}
			g.screen.DrawAt(strconv.Itoa(idx)+":"+card.Ansi(), i, 8+idx)
		}

		i += width
		c += 1
	}

	in := fmt.Sprintf("Input: %s", g.InputBuf.String())

	g.screen.DrawAt(in, 0, 20)

	g.screen.Flush()

	time.Sleep(250 * time.Millisecond)
}
