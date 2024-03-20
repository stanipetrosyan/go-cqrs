package main

func main() {
	println("Bank account cqrs example")
	commandBus := NewCommandBus()

	createAccount := CreateAccount{name: "Stani"}

	commandBus.apply(createAccount)

	depositMoney := DepositMoney{value: 10}

	commandBus.apply(depositMoney)

}

type Command interface{}

type CreateAccount struct {
	name string
}

type DepositMoney struct {
	value int
}
