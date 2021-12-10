package banking

import "errors"

var ErrAccountInsufficientFunds = errors.New("account has insufficient funds")
var ErrATMInsufficientFunds = errors.New("ATM has insufficient funds")
var ErrCardRetained = errors.New("card has been retained")

type ATM struct {
	card  *Card
	funds int
}

func NewATM(money int) *ATM {
	return &ATM{funds: money}
}

func (a *ATM) Withdraw(c *Cardholder, amount int) (int, error) {
	switch {
	case c == nil:
		return 0, errors.New("cardholder must not be nil")
	case c.card == nil:
		return 0, errors.New("card must not be nil")
	case c.card.account == nil:
		return 0, errors.New("account must not be nil")
	case c.card.Disabled():
		c.card = nil
		return 0, ErrCardRetained
	case !c.card.Valid():
		return 0, errors.New("card must be valid")
	case amount > c.card.account.funds:
		return 0, ErrAccountInsufficientFunds
	case amount > a.funds:
		return 0, ErrATMInsufficientFunds
	}

	c.card.account.funds -= amount

	return amount, nil
}
