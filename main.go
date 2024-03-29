package main

import (
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

func main() {
	println("Bank account cqrs example")
	eventbus := goeventbus.NewEventBus()
	eventstore := NewEventStore(eventbus)
	commandBus := NewCommandBus(eventstore)

	accountView := NewAccountView(eventbus)
	accountView.Listen()

	createAccount := CreateAccount{name: "user"}

	commandBus.apply(createAccount)

	depositMoney := DepositMoney{name: "user", value: 10}

	commandBus.apply(depositMoney)

	withdrawMoney := WithdrawMoney{name: "user", value: 20}

	commandBus.apply(withdrawMoney)

	println(accountView.GetAccounts()[0])
}

type View interface {
	Listen()
}

type AccountsView struct {
	eventbus goeventbus.EventBus
	accounts []string
}

func (v *AccountsView) Listen() {
	v.eventbus.Channel("AccountCreated").Subscriber().Listen(func(context goeventbus.Context) {
		println("account created:", context.Result().Data.(AccountCreated).name)
		v.accounts = append(v.accounts, context.Result().Data.(AccountCreated).name)
	})
}

func (v *AccountsView) GetAccounts() []string {
	return v.accounts
}

func NewAccountView(eventbus goeventbus.EventBus) *AccountsView {
	return &AccountsView{eventbus: eventbus, accounts: []string{}}
}
