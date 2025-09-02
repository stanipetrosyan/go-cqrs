package main

import (
	"sync"

	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	println("Bank account cqrs example")
	eventbus := goeventbus.NewEventBus()
	eventstore := NewEventStore(eventbus)
	commandBus := NewCommandBus(eventstore)

	//
	NewWorkflow(eventbus, eventstore).Step(SagaStep{Transaction: "MoneyDeposited"}).Step(SagaStep{Transaction: "MoneyWithdrawn"}).Commit()
	//
	accountProjection := NewAccountProjection(eventbus)
	accountProjection.Listen()

	createAccount := CreateAccount{name: "user"}

	commandBus.apply(createAccount)

	/* 	println("Accounts:", accountProjection.GetAccounts()[0])
	 */
	println("Depositing 10 dollars")
	depositMoney := DepositMoney{name: "user", value: 10}

	commandBus.apply(depositMoney)

	println("Withdrawing 5 dollars")
	withdrawMoney := WithdrawMoney{name: "user", value: 5}

	commandBus.apply(withdrawMoney)
	wg.Wait()
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
		v.accounts = append(v.accounts, context.Result().Extract().(AccountCreated).name)
	})
}

func (v *AccountsProjection) GetAccounts() []string {
	return v.accounts
}

func NewAccountProjection(eventbus goeventbus.EventBus) *AccountsProjection {
	return &AccountsProjection{eventbus: eventbus, accounts: []string{}}
}
