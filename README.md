# scenario

scenario is an extremely simple BDD library (~100 LOC) for Golang that's 100% compatible with the standard testing package.

## How To Use

```go
func TestATM_Withdraw(t *testing.T) {
	t.Parallel()
    
	// using scenario.Title is optional: you can start with s := scenario.Given(...
	s := scenario.Title("Account has insufficient funds").
		Given("the account funds is $100").
		And("the card is valid").
		And("the ATM contains enough funds").
		When("the Cardholder requests $20")

	funds := 10
	request := 20
	wantDispensed := 0
	wantFunds := 10

	account := banking.NewAccount(funds)

	card := banking.NewCard(account, false)
	cardholder := banking.NewCardholder(card)

	atm := banking.NewATM(request)

	dispensed, err := atm.Withdraw(cardholder, request)

	s.Then("the ATM should dispense $20", func(t *testing.T) {
		t.Parallel()

		assert.ErrorIs(t, err, banking.ErrAccountInsufficientFunds)
		assert.Equal(t, wantDispensed, dispensed)
	})

	s.And("the account funds should be $80", func(t *testing.T) {
		t.Parallel()

		assert.Equal(t, wantFunds, account.Funds())
	})

	s.And("the card should be returned", func(t *testing.T) {
		t.Parallel()

		assert.NotNil(t, cardholder.Card())
	})

	s.Run(t)
}
```
