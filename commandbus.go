package main

type CommandBus interface {
	apply(command Command)
}

type defaultCommandBus struct {
	eventstore EventStore
}

func (c defaultCommandBus) apply(command Command) {
	switch command := command.(type) {
	case CreateAccount:
		CreateAccountHandler{eventstore: c.eventstore}.handle(command)
	case DepositMoney:
		DepositMoneyHandler{eventstore: c.eventstore}.handle(command)
	case WithdrawMoney:
		WithdrawMoneyHandler{eventstore: c.eventstore}.handle(command)
	}
}

func NewCommandBus(eventstore EventStore) CommandBus {
	return defaultCommandBus{eventstore: eventstore}
}
