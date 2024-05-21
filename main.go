package main

import (
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

func main() {
	println("Bank account cqrs example")
	eventbus := goeventbus.NewEventBus()
	eventstore := NewEventStore(eventbus)
	commandBus := NewCommandBus(eventstore)

	accountProjection := NewAccountProjection(eventbus)
	accountProjection.Listen()

	createAccount := CreateAccount{name: "user"}

	commandBus.apply(createAccount)

	println("Accounts:", accountProjection.GetAccounts()[0])

	println("Depositing 10 dollars")
	depositMoney := DepositMoney{name: "user", value: 10}

	commandBus.apply(depositMoney)

	println("Withdrawing 20 dollars")
	withdrawMoney := WithdrawMoney{name: "user", value: 20}

	commandBus.apply(withdrawMoney)

}

type Projection interface {
	Listen()
}

type AccountsProjection struct {
	eventbus goeventbus.EventBus
	accounts []string
}

func (v *AccountsProjection) Listen() {
	v.eventbus.Channel("AccountCreated").Subscriber().Listen(func(context goeventbus.Context) {
		v.accounts = append(v.accounts, context.Result().Data.(AccountCreated).name)
	})
}

func (v *AccountsProjection) GetAccounts() []string {
	return v.accounts
}

func NewAccountProjection(eventbus goeventbus.EventBus) *AccountsProjection {
	return &AccountsProjection{eventbus: eventbus, accounts: []string{}}
}
