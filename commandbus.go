package main

type CommandBus interface {
	apply(command Command)
}

type defaultCommandBus struct {
}

func (c defaultCommandBus) apply(command Command) {
	switch command := command.(type) {
	case CreateAccount:
		CreateAccountHandler{}.handle(command)
	case DepositMoney:
		DepositMoneyHandler{}.handle(command)

	}
}

func NewCommandBus() CommandBus {
	return defaultCommandBus{}
}

type CommandHandler[T Command] interface {
	handle(command T)
}

type CreateAccountHandler struct{}

func (h CreateAccountHandler) handle(command CreateAccount) {
	println(command.name)
}

type DepositMoneyHandler struct{}

func (h DepositMoneyHandler) handle(command DepositMoney) {
	println(command.value)
}
