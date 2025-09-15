package main

import (
	"fmt"
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
	err := NewWorkflow(eventbus, eventstore).
		EntryPoint(SagaStep{Transaction: "WireTransferStarted"}).
		Step(SagaStep{Transaction: "MoneyDeposited"}).
		Step(SagaStep{Transaction: "MoneyWithdrawn", Compensate: "MoneyWithdrawnRejected"}).
		Commit()
	//

	if err != nil {
		fmt.Println(err)
	}
	accountProjection := NewAccountProjection(eventbus)
	accountProjection.Listen()

	createAccount := CreateAccount{name: "user"}

	commandBus.apply(createAccount)

	transactionId := "randomUUID"
	wireTransferStart := WireTransferStart{transactionId: transactionId, name: "user"}
	commandBus.apply(wireTransferStart)

	println("Depositing 10 dollars")
	depositMoney := DepositMoney{transactionId: transactionId, name: "user", value: 10}

	commandBus.apply(depositMoney)

	println("Withdrawing 5 dollars")
	withdrawMoney := WithdrawMoney{transactionId: transactionId, name: "user", value: 5}

	commandBus.apply(withdrawMoney)

	/* 	createAccount = CreateAccount{name: "anotherUser"}

	   	commandBus.apply(createAccount)
	   	transactionId = "anotherRandomUUID"
	   	wireTransferStart = WireTransferStart{transactionId: transactionId, name: "anotherUser"}
	   	commandBus.apply(wireTransferStart)

	   	println("Depositing 10 dollars")
	   	depositMoney = DepositMoney{transactionId: transactionId, name: "anotherUser", value: 10}

	   	commandBus.apply(depositMoney)

	   	println("Withdrawing 5 dollars")
	   	withdrawMoney = WithdrawMoney{transactionId: transactionId, name: "anotherUser", value: 20}

	   	commandBus.apply(withdrawMoney) */

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
