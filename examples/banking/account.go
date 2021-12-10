package banking

type Account struct {
	funds int
}

func NewAccount(funds int) *Account {
	return &Account{funds: funds}
}

func (a *Account) Funds() int {
	return a.funds
}
