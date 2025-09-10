package main

type Event interface {
	Aggregate() string
	eventName() string
}

type TransactionEvent interface {
	Event
	Transaction() string
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
	transactionId string
	name          string
	value         int
}

func (e MoneyDeposited) eventName() string {
	return "MoneyDeposited"
}

func (e MoneyDeposited) Aggregate() string {
	return e.name
}

func (e MoneyDeposited) Transaction() string {
	return e.transactionId
}

type MoneyWithdrawn struct {
	transactionId string
	name          string
	value         int
}

func (e MoneyWithdrawn) eventName() string {
	return "MoneyWithdrawn"
}

func (e MoneyWithdrawn) Aggregate() string {
	return e.name
}

func (e MoneyWithdrawn) Transaction() string {
	return e.transactionId
}

type MoneyWithdrawnRejected struct {
	transactionId string
	name          string
	value         int
}

func (e MoneyWithdrawnRejected) eventName() string {
	return "MoneyWithdrawnRejected"
}

func (e MoneyWithdrawnRejected) Aggregate() string {
	return e.name
}

func (e MoneyWithdrawnRejected) Transaction() string {
	return e.transactionId
}

type WireTransferStarted struct {
	transactionId string
	name          string
}

func (e WireTransferStarted) eventName() string {
	return "WireTransferStarted"
}

func (e WireTransferStarted) Aggregate() string {
	return e.name
}

func (e WireTransferStarted) Transaction() string {
	return e.transactionId
}

type WireTransferCompleted struct {
	transaction string
}

func (e WireTransferCompleted) eventName() string {
	return "WireTransferCompleted"
}

func (e WireTransferCompleted) Aggregate() string {
	return e.transaction
}

type WireTransferRejected struct {
	transaction string
}

func (e WireTransferRejected) eventName() string {
	return "WireTransferRejected"
}

func (e WireTransferRejected) Aggregate() string {
	return e.transaction
}
