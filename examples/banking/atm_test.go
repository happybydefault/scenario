package banking_test

import (
	"fmt"
	"testing"

	"github.com/happybydefault/scenario"
	"github.com/happybydefault/scenario/examples/banking"

	"github.com/stretchr/testify/assert"
)

func TestATM_Withdraw(t *testing.T) {
	testATMWithdrawAccountHasInsufficientFunds(t)
	testATMWithdrawCardHasBeenDisabled(t)
	testATMWithdrawATMHasInsufficientFunds(t)
}

func testATMWithdrawAccountHasInsufficientFunds(t *testing.T) {
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

	s.And("the account balance should be the same as the beginning", func(t *testing.T) {
		assert.Equal(t, funds, account.Funds())
	})

	s.And("the card should be returned", func(t *testing.T) {
		assert.NotNil(t, cardholder.Card())
	})

	s.Run(t)
}

func testATMWithdrawCardHasBeenDisabled(t *testing.T) {
	s := scenario.New("Card has been disabled").
		Given("the card is disabled").
		When("the Account Holder requests $20")

	cardInvalid := true
	request := 20

	account := banking.NewAccount(request)

	card := banking.NewCard(account, cardInvalid)
	cardholder := banking.NewCardholder(card)

	atm := banking.NewATM(request)

	_, err := atm.Withdraw(cardholder, request)

	s.Then("the ATM should retain the card", func(t *testing.T) {
		assert.Nil(t, cardholder.Card())
	})

	s.And("the ATM should say the card has been retained", func(t *testing.T) {
		assert.ErrorIs(t, err, banking.ErrCardRetained)
	})

	s.Run(t)
}

func testATMWithdrawATMHasInsufficientFunds(t *testing.T) {
	type testCase struct {
		funds     int
		request   int
		wantFunds int
	}

	testCases := []testCase{
		{
			funds:     100,
			request:   20,
			wantFunds: 100,
		},
		{
			funds:     150,
			request:   20,
			wantFunds: 150,
		},
		{
			funds:     200,
			request:   30,
			wantFunds: 200,
		},
	}

	for _, tc := range testCases {
		s := scenario.New("ATM has insufficient funds").
			Given(fmt.Sprintf("the account balance is $%d", tc.funds)).
			And("the card is valid").
			And("the machine does not contain enough funds").
			When(fmt.Sprintf("the Account Holder requests $%d", tc.request))

		account := banking.NewAccount(tc.funds)

		card := banking.NewCard(account, false)
		cardholder := banking.NewCardholder(card)

		atm := banking.NewATM(tc.request - 1)

		dispensed, err := atm.Withdraw(cardholder, tc.request)

		s.Then("the ATM should say it has insufficient funds", func(t *testing.T) {
			assert.ErrorIs(t, err, banking.ErrATMInsufficientFunds)
		})

		s.And("the ATM should not dispense any funds", func(t *testing.T) {
			assert.Equal(t, 0, dispensed)
		})

		s.And(fmt.Sprintf("the account balance should be $%d", tc.wantFunds), func(t *testing.T) {
			assert.Equal(t, tc.wantFunds, account.Funds())
		})

		s.And("the card should be returned", func(t *testing.T) {
			assert.NotNil(t, cardholder.Card(), "the card should be returned")
		})

		s.Run(t)
	}
}
