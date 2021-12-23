# scenario

scenario is an extremely simple BDD library (~100 LOC) for Golang that's 100% compatible with the standard testing package.

## How To Use

```go
func TestATM_Withdraw(t *testing.T) {
	s := scenario.New("Account has insufficient funds").
		Given("the account funds is $100").
		And("the card is valid").
		And("the ATM contains enough funds").
		When("the Cardholder requests $20")

	funds := 100
	request := 200
	cardInvalid := false

	account := banking.NewAccount(funds)

	card := banking.NewCard(account, cardInvalid)
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

### Output

```sh
go test ./examples/banking -run "^TestATM_Withdraw$" -v -scenario.pretty
```

```
=== RUN   TestATM_Withdraw
Scenario: Account has insufficient funds
Given the account funds is $100
And the card is valid
And the ATM contains enough funds
When the Cardholder requests $20
Then the ATM should dispense $0
And the account funds should be $100
And the card should be returned

=== RUN   TestATM_Withdraw/Account_has_insufficient_funds
=== RUN   TestATM_Withdraw/Account_has_insufficient_funds/the_ATM_should_dispense_$0
=== RUN   TestATM_Withdraw/Account_has_insufficient_funds/the_account_funds_should_be_$100
=== RUN   TestATM_Withdraw/Account_has_insufficient_funds/the_card_should_be_returned
--- PASS: TestATM_Withdraw (0.00s)
    --- PASS: TestATM_Withdraw/Account_has_insufficient_funds (0.00s)
PASS
ok      scenario/examples/banking       0.002s
```
