package main

type Event interface {
	eventName() string
}

type AccountCreated struct {
	name string
}

func (e AccountCreated) eventName() string {
	return "AccountCreated"
}

type MoneyDeposited struct {
	name  string
	value int
}

func (e MoneyDeposited) eventName() string {
	return "MoneyDeposited"
}

type MoneyWithdrawn struct {
	name  string
	value int
}

func (e MoneyWithdrawn) eventName() string {
	return "MoneyWithdrawn"
}
