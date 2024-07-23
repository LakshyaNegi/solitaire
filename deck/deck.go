package deck

import "math/rand"

type Value string

const (
	Ace   Value = "A"
	Two   Value = "2"
	Three Value = "3"
	Four  Value = "4"
	Five  Value = "5"
	Six   Value = "6"
	Seven Value = "7"
	Eight Value = "8"
	Nine  Value = "9"
	Ten   Value = "10"
	Jack  Value = "J"
	Queen Value = "Q"
	King  Value = "K"
)

var Values [13]Value = [13]Value{
	Ace,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Ten,
	Jack,
	Queen,
	King,
}

type Deck struct {
	Cards []*Card
	Total int
}

func (d *Deck) GetTop() *Card {
	if len(d.Cards) == 0 {
		return nil
	}

	return d.Cards[len(d.Cards)-1]
}

func (d *Deck) GetIdx(idx int) *Card {
	if len(d.Cards) <= idx {
		return nil
	}

	return d.Cards[idx]
}

func (d *Deck) PopCard() *Card {
	if d.Total == 0 {
		return nil
	}
	card := d.Cards[d.Total-1]

	d.Total--
	d.Cards = d.Cards[:d.Total]

	return card
}

func (d *Deck) PutCard(card *Card) {
	d.Cards = append(d.Cards, card)
	d.Total++
}

func NewFromCards(cards []*Card) Deck {
	return Deck{
		Cards: cards,
		Total: len(cards),
	}
}

func NewDeck() Deck {
	cards := []*Card{}
	for i := range 4 {
		for j := range 13 {
			cards = append(cards, &Card{
				Suite: Suites[i],
				Value: Values[j],
			})
		}
	}

	for i := range 52 {
		idx := rand.Intn(52 - i)
		last := cards[len(cards)-1-i]
		cards[len(cards)-1-i] = cards[idx]
		cards[idx] = last
	}

	return Deck{
		Cards: cards,
		Total: 52,
	}
}
