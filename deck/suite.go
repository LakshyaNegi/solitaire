package deck

const (
	Diamonds Suite = "D"
	Hearts   Suite = "H"
	Clubs    Suite = "C"
	Spades   Suite = "S"
)

var BlackSuite = map[Suite]bool{Spades: true, Clubs: true}
var RedSuite = map[Suite]bool{Hearts: true, Diamonds: true}

// ANSI escape codes for colors
const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[31m"
	ColorBlack = "\033[34m"
)

// Unicode characters for card suits
const (
	ANSIHeart   = "\u2665"
	ANSIDiamond = "\u2666"
	ANSIClub    = "\u2663"
	ANSISpade   = "\u2660"
)

type Suite string

func (s Suite) Ansi() string {
	switch s {
	case Diamonds:
		return ColorRed + ANSIDiamond + ColorReset
	case Hearts:
		return ColorRed + ANSIHeart + ColorReset
	case Spades:
		return ColorBlack + ANSISpade + ColorReset
	case Clubs:
		return ColorBlack + ANSIClub + ColorReset
	}

	return ""
}

var Suites [4]Suite = [4]Suite{Diamonds, Hearts, Clubs, Spades}
