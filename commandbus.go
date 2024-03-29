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

type CommandHandler interface {
	handle(command Command)
}

type CreateAccountHandler struct {
	eventstore EventStore
}

func (h CreateAccountHandler) handle(command CreateAccount) {
	HydrateBankAccount(h.eventstore.load(command.name))
	event := AccountCreated{name: command.name}

	h.eventstore.save(command.name, event)
}

type DepositMoneyHandler struct {
	eventstore EventStore
}

func (h DepositMoneyHandler) handle(command DepositMoney) {
	HydrateBankAccount(h.eventstore.load(command.name))
	event := MoneyDeposited{name: command.name, value: command.value}

	h.eventstore.save(command.name, event)
}

type WithdrawMoneyHandler struct {
	eventstore EventStore
}

func (h WithdrawMoneyHandler) handle(command WithdrawMoney) {
	account := HydrateBankAccount(h.eventstore.load(command.name))

	if account.CanWithdrawn(command.value) {
		event := MoneyWithdrawn{name: command.name, value: command.value}

		h.eventstore.save(command.name, event)
	} else {
		println("cannot perform command")
	}

}
