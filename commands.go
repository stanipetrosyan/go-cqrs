package main

type Command interface{}

type CreateAccount struct {
	name string
}

type WithdrawMoney struct {
	name  string
	value int
}

type DepositMoney struct {
	name  string
	value int
}
