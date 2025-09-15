package main

type Event interface {
	Aggregate() string
	EventName() string
}

type TransactionEvent interface {
	Event
	Transaction() string
}

type AccountCreated struct {
	name string
}

func (e AccountCreated) EventName() string {
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

func (e MoneyDeposited) EventName() string {
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

func (e MoneyWithdrawn) EventName() string {
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

func (e MoneyWithdrawnRejected) EventName() string {
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

func (e WireTransferStarted) EventName() string {
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

func (e WireTransferCompleted) EventName() string {
	return "WireTransferCompleted"
}

func (e WireTransferCompleted) Aggregate() string {
	return e.transaction
}

type WireTransferRejected struct {
	transaction string
}

func (e WireTransferRejected) EventName() string {
	return "WireTransferRejected"
}

func (e WireTransferRejected) Aggregate() string {
	return e.transaction
}
