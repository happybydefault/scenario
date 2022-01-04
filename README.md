# scenario

scenario is an extremely simple BDD library (~100 sloc) for Golang that's 100% compatible with the standard testing
package.

## How To Use

```go
import (
	"testing"

	"github.com/happybydefault/scenario"
	"github.com/happybydefault/scenario/examples/banking"

	"github.com/stretchr/testify/assert"
)

func TestATM_Withdraw_AccountHasInsufficientFunds(t *testing.T) {
	s := scenario.New("Account has insufficient funds").
		Given("the account balance is $10").
		And("the card is valid").
		And("the machine contains enough funds").
		When("the Account Holder requests $20")

	funds := 10
	request := 20
	cardInvalid := false

	account := banking.NewAccount(funds)

	card := banking.NewCard(account, cardInvalid)
	cardholder := banking.NewCardholder(card)

	atm := banking.NewATM(request)

	dispensed, err := atm.Withdraw(cardholder, request)

	s.Then("the ATM should not dispense any money", func(t *testing.T) {
		assert.Equal(t, 0, dispensed)
	})

	s.And("the ATM should say there are insufficient funds", func(t *testing.T) {
		assert.ErrorIs(t, err, banking.ErrAccountInsufficientFunds)
	})

	s.And("the account balance should be the same as initially", func(t *testing.T) {
		assert.Equal(t, funds, account.Funds())
	})

	s.And("the card should be returned", func(t *testing.T) {
		assert.NotNil(t, cardholder.Card())
	})

	s.Run(t)
}
```

### Output

#### Normal

```sh
go test ./examples/banking -run "TestATM_Withdraw_AccountHasInsufficientFunds" -v
```

```
=== RUN   TestATM_Withdraw_AccountHasInsufficientFunds
    atm_test.go:49: Scenario: Account has insufficient funds
        Given the account balance is $10
        And the card is valid
        And the machine contains enough funds
        When the Account Holder requests $20
        Then the ATM should not dispense any money
        And the ATM should say there are insufficient funds
        And the account balance should be the same as initially
        And the card should be returned
=== RUN   TestATM_Withdraw_AccountHasInsufficientFunds/Account_has_insufficient_funds
=== RUN   TestATM_Withdraw_AccountHasInsufficientFunds/Account_has_insufficient_funds/the_ATM_should_not_dispense_any_money
=== RUN   TestATM_Withdraw_AccountHasInsufficientFunds/Account_has_insufficient_funds/the_ATM_should_say_there_are_insufficient_funds
=== RUN   TestATM_Withdraw_AccountHasInsufficientFunds/Account_has_insufficient_funds/the_account_balance_should_be_the_same_as_initially
=== RUN   TestATM_Withdraw_AccountHasInsufficientFunds/Account_has_insufficient_funds/the_card_should_be_returned
--- PASS: TestATM_Withdraw_AccountHasInsufficientFunds (0.00s)
    --- PASS: TestATM_Withdraw_AccountHasInsufficientFunds/Account_has_insufficient_funds (0.00s)
        --- PASS: TestATM_Withdraw_AccountHasInsufficientFunds/Account_has_insufficient_funds/the_ATM_should_not_dispense_any_money (0.00s)
        --- PASS: TestATM_Withdraw_AccountHasInsufficientFunds/Account_has_insufficient_funds/the_ATM_should_say_there_are_insufficient_funds (0.00s)
        --- PASS: TestATM_Withdraw_AccountHasInsufficientFunds/Account_has_insufficient_funds/the_account_balance_should_be_the_same_as_initially (0.00s)
        --- PASS: TestATM_Withdraw_AccountHasInsufficientFunds/Account_has_insufficient_funds/the_card_should_be_returned (0.00s)
PASS
ok  	github.com/happybydefault/scenario/examples/banking	(cached)

```

#### With flag `-scenario.pretty`

The scenario description is colored light blue.

```sh
go test ./examples/banking -run "TestATM_Withdraw_AccountHasInsufficientFunds" -v -scenario.pretty
```

![Output with flag scenario dot pretty](assets/pretty.png "Output with flag -scenario.pretty")

#### With flag `-scenario.pretty` and program `prettytest` (recommended)

The scenario description is colored light blue, and the whole output from `go test` is processed
by [prettytest](https://github.com/happybydefault/prettytest).

```sh
prettytest ./examples/banking -run "^TestATM_Withdraw$" -v -scenario.pretty
```

![Output with flag scenario dot pretty and program pretty test](assets/prettytest.png "Output with flag -scenario.pretty and program prettytest")
