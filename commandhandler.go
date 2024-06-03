package main

type CommandHandler interface {
	handle(command Command)
}

type CreateAccountHandler struct {
	eventstore EventStore
}

func (h CreateAccountHandler) handle(command CreateAccount) {
	HydrateBankAccount(h.eventstore.Load(command.name))
	event := AccountCreated{name: command.name}

	h.eventstore.Save(command.name, event)
}

type DepositMoneyHandler struct {
	eventstore EventStore
}

func (h DepositMoneyHandler) handle(command DepositMoney) {
	HydrateBankAccount(h.eventstore.Load(command.name))
	event := MoneyDeposited{name: command.name, value: command.value}

	h.eventstore.Save(command.name, event)
}

type WithdrawMoneyHandler struct {
	eventstore EventStore
}

func (h WithdrawMoneyHandler) handle(command WithdrawMoney) {
	account := HydrateBankAccount(h.eventstore.Load(command.name))

	if account.CanWithdrawn(command.value) {
		event := MoneyWithdrawn{name: command.name, value: command.value}

		h.eventstore.Save(command.name, event)
	} else {
		println("Cannot perform command: WithdrawMoney")
	}

}
