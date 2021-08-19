package model

import "time"

const (
	CardTypeVirtual = "VIRTUAL"
)

type Card struct {
	ID uint `json:"id"`
	Type string `json:"type"`
	Created time.Time `json:"created"`
	CustomerID uint
	Customer Customer
	AccountID uint
	Account Account
}
