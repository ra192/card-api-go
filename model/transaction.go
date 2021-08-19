package model

import "time"

const (
	TransactionTypeFund="FUND"
	TransactionTypeVirtualCardDeposit = "VIRTUAL_CARD_DEPOSIT"
	TransactionTypeVirtualCardWithdraw = "VIRTUAL_CARD_WITHDRAW"
)

type Transaction struct {
	ID uint
	OrderId string
	Type string
	Status string
}

type TransactionItem struct {
	ID uint
	Amount int
	Created time.Time
	SourceAccountID uint
	SourceAccount Account
	DestinationAccountID uint
	DestinationAccount Account
	CardID uint
	Card Card
}

type TransactionFee struct {
	ID uint
	rate float32
	Type string
	AccountID uint
	Account Account
}