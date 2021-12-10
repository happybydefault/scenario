package banking

type Cardholder struct {
	card *Card
}

func NewCardholder(card *Card) *Cardholder {
	return &Cardholder{card: card}
}

func (c *Cardholder) Card() *Card {
	return c.card
}
