package banking_test

import (
	"fmt"
	"testing"

	"github.com/happybydefault/scenario"
	"github.com/happybydefault/scenario/examples/banking"

	"github.com/stretchr/testify/assert"
)

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

func TestATM_Withdraw_AccountHasInsufficientFunds(t *testing.T) {
	s := scenario.New("Account has insufficient funds").
		Given("the account balance is $10").
		And("the card is valid").
		And("the machine contains enough funds").
		When("the Account Holder requests $20")

	funds := 10
	request := 20
	wantDispensed := 0
	wantFunds := 10

	account := banking.NewAccount(funds)

	card := banking.NewCard(account, false)
	cardholder := banking.NewCardholder(card)

	atm := banking.NewATM(request)

	dispensed, err := atm.Withdraw(cardholder, request)

	s.Then("the ATM should not dispense any funds", func(t *testing.T) {
		assert.ErrorIs(t, err, banking.ErrAccountInsufficientFunds)
		assert.Equal(t, wantDispensed, dispensed)
	})

	s.And("the ATM should say there are insufficient funds", func(t *testing.T) {
		assert.Equal(t, wantFunds, account.Funds())
	})

	s.And("the card should be returned", func(t *testing.T) {
		assert.NotNil(t, cardholder.Card())
	})

	s.Run(t)
}

func TestATM_Withdraw_CardHasBeenDisabled(t *testing.T) {
	s := scenario.New("Card has been disabled").
		Given("the card is disabled").
		When("the Account Holder requests $20")

	request := 20

	account := banking.NewAccount(request)

	card := banking.NewCard(account, true)
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

func TestATM_Withdraw_ATMHasInsufficientFunds(t *testing.T) {
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
