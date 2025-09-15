package main

type Command interface{}

type CreateAccount struct {
	name string
}

type WithdrawMoney struct {
	transactionId string
	name          string
	value         int
}

type DepositMoney struct {
	transactionId string
	name          string
	value         int
}

type WireTransferStart struct {
	transactionId string
	name          string
}
