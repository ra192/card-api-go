package model

import "time"

const (
	CardTypeVirtual = "VIRTUAL"
)

type Card struct {
	ID         uint
	Type       string
	Created    time.Time
	CustomerID uint
	Customer   Customer
	AccountID  uint
	Account    Account
}
