# scenario

scenario is an extremely simple BDD library (~100 LOC) for Golang that's 100% compatible with the standard testing package.

## How To Use

```go
func TestATM_Withdraw(t *testing.T) {
	// using scenario.Title is optional: you can start with s := scenario.Given(...
	s := scenario.Title("Account has insufficient funds").
		Given("the account funds is $100").
		And("the card is valid").
		And("the ATM contains enough funds").
		When("the Cardholder requests $200")

	funds := 100
	validCard := true
	request := 200

	account := banking.NewAccount(funds)

	card := banking.NewCard(account, validCard)
	cardholder := banking.NewCardholder(card)

	atm := banking.NewATM(request)

	dispensed, err := atm.Withdraw(cardholder, request)

	s.Then("the ATM should dispense $0", func(t *testing.T) {
		assert.ErrorIs(t, err, banking.ErrAccountInsufficientFunds)
		assert.Equal(t, 0, dispensed)
	})

	s.And("the account funds should be $100", func(t *testing.T) {
		assert.Equal(t, 100, account.Funds())
	})

	s.And("the card should be returned", func(t *testing.T) {
		assert.NotNil(t, cardholder.Card())
	})

	s.Run(t)
}
```
