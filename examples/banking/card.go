package banking

type Card struct {
	account  *Account
	disabled bool
}

func NewCard(account *Account, disabled bool) *Card {
	return &Card{account: account, disabled: disabled}
}

// Valid differs from Disabled, as Valid could also check for Card expiration and other stuff.
func (c *Card) Valid() bool {
	return !c.disabled
}

func (c *Card) Disable() {
	c.disabled = true
}

func (c *Card) Disabled() bool {
	return c.disabled
}
