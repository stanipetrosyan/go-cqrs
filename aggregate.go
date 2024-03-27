package main

type BankAccount struct {
	name  string
	money int
}

func (a BankAccount) CanWithdrawn(value int) bool {
	return a.money >= value
}

func HydrateBankAccount(events []Event) BankAccount {
	account := BankAccount{}
	for _, event := range events {
		switch event := event.(type) {
		case AccountCreated:
			account.name = event.name
		case MoneyDeposited:
			account.money += event.value
		case MoneyWithdrawn:
			account.money -= event.value
		}
	}

	return account
}
