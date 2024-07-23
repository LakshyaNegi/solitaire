package deck

import (
	"log"
	"solitaire/screen"
	"strconv"
)

type Card struct {
	Suite  Suite
	Value  Value
	IsOpen bool
}

func (c *Card) Ansi() string {
	if c == nil {
		return "00"
	}
	if !c.IsOpen {
		return "##"
	}

	return c.Suite.Ansi() + string(c.Value)
}

func (c *Card) GetNumericValue() int {
	if c == nil {
		return 0
	}

	switch c.Value {
	case King:
		return 13
	case Queen:
		return 12
	case Jack:
		return 11
	case Ace:
		return 1
	default:
		i, err := strconv.Atoi(string(c.Value))
		if err != nil {
			log.Fatal("card value not determinable")
		}

		return i
	}
}

func (c *Card) DrawAt(screen screen.Screen, x, y int) {
	if c == nil {
		screen.DrawAt("00", x, y)
		return
	}

	if !c.IsOpen {
		screen.DrawAt("##", x, y)
		return
	}

	screen.DrawAt(c.Suite.Ansi()+string(c.Value), x, y)
}
