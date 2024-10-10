package main

type Event interface {
	Aggregate() string
	eventName() string
}

type AccountCreated struct {
	name string
}

func (e AccountCreated) eventName() string {
	return "AccountCreated"
}

func (e AccountCreated) Aggregate() string {
	return e.name
}

type MoneyDeposited struct {
	name  string
	value int
}

func (e MoneyDeposited) eventName() string {
	return "MoneyDeposited"
}

func (e MoneyDeposited) Aggregate() string {
	return e.name
}

type MoneyWithdrawn struct {
	name  string
	value int
}

func (e MoneyWithdrawn) eventName() string {
	return "MoneyWithdrawn"
}

func (e MoneyWithdrawn) Aggregate() string {
	return e.name
}
